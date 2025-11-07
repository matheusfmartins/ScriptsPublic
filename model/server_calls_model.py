from database import db
from flask import jsonify
from datetime import datetime, timedelta

# Check all agents
def get_agents():
    conn = db.get_conn()
    cursor = conn.cursor()

    cursor.execute("""
        SELECT agent_id, agent_username, agent_hostname, agent_os, last_seen
        FROM agents
    """)
    results = cursor.fetchall()
    conn.close()

    agents = []
    
    for row in results:
        
        last_seen = humanize_time_diff(row[4])
        
        agent = {
            "agent_id": row[0],
            "agent_username": row[1],
            "agent_hostname": row[2],
            "agent_os": row[3],
            "last_seen": last_seen
        }
        agents.append(agent)
    
    return jsonify({"agents": agents})

# Check agent by id
def get_agent_by_id(agent_id):
    conn = db.get_conn()
    cursor = conn.cursor()

    cursor.execute("""
        SELECT agent_id, agent_username, agent_hostname, agent_os, last_seen
        FROM agents 
        WHERE agent_id = ?
    """, (agent_id,))
    result = cursor.fetchall()
    conn.close()

    result = result[0]

    last_seen = humanize_time_diff(result[4])
        
    agent = {
        "agent_id": result[0],
        "agent_username": result[1],
        "agent_hostname": result[2],
        "agent_os": result[3],
        "last_seen": last_seen
    }
    
    return jsonify(agent)

# Create a new command for a agent
def create_command(agent_id, agent_command):
    conn = db.get_conn()
    cursor = conn.cursor()
    print("criando")
    command_datetime = datetime.now()
    
    cursor.execute("""
        INSERT INTO commands (agent_id, agent_command, command_datetime)
        VALUES (?, ?, ?)
    """, (agent_id, agent_command, command_datetime))
    
    conn.commit()
    conn.close()
    
    return jsonify({"result": "ok"})

# Create a new command for a agent
def get_commands(agent_id):
    conn = db.get_conn()
    cursor = conn.cursor()
    
    # search commands
    cursor.execute("""
        SELECT command_id, agent_id, agent_command, command_result, command_datetime
        FROM commands
        WHERE agent_id = ?
        ORDER BY agent_id
    """, (agent_id,))
    commands = cursor.fetchall()
    
    conn.close()
    
    processed_commands = []
    
    for command in commands:
        processed_command = {
            "command_id": command[0],
            "agent_id": command[1],
            "agent_command": command[2],
            "command_result": command[3],
            "command_datetime": command[4]
        }
        
        processed_commands.append(processed_command)
    
    return jsonify({"commands": processed_commands})

# Delete inactive agents
def delete_inactive_agents():
    conn = db.get_conn()
    cursor = conn.cursor()
    try:
        # Hora atual e limite de inatividade (1 minuto atrás)
        now = datetime.now()
        inactivity_threshold = now - timedelta(minutes=1)

        # Buscar agentes inativos
        cursor.execute("SELECT agent_id, last_seen FROM agents")
        all_agents = cursor.fetchall()

        inactive_agents = []
        for agent_id, last_seen_str in all_agents:
            try:
                last_seen = datetime.strptime(last_seen_str, "%Y-%m-%d %H:%M:%S.%f")
            except ValueError:
                continue
            
            if last_seen < inactivity_threshold:
                
                inactive_agents.append(agent_id)

        # Deletar agentes inativos
        for agent_id in inactive_agents:
            cursor.execute("DELETE FROM agents WHERE agent_id = ?", (agent_id,))
            cursor.execute("DELETE FROM commands WHERE agent_id = ?", (agent_id,))

        conn.commit()
        conn.close()

        return jsonify({
            "status": "success",
            "deleted_count": len(inactive_agents),
            "deleted_agents": inactive_agents
        }), 200

    except Exception as e:
        return jsonify({"status": "error", "message": str(e)}), 500

def humanize_time_diff(past):
    now = datetime.now()
    
    if isinstance(past, str):
        past = datetime.fromisoformat(past)

    diff = now - past
    seconds = int(diff.total_seconds())

    exact_time = past.strftime("%d/%m/%Y %H:%M")

    if seconds < 60:
        return f"{seconds} second{'s' if seconds != 1 else ''} ago ({exact_time})"
    elif seconds < 3600:
        minutes = int(seconds / 60)
        return f"{minutes} minute{'s' if minutes != 1 else ''} ago ({exact_time})"
    elif seconds < 86400:
        hours = int(seconds / 3600)
        return f"{hours} hour{'s' if hours != 1 else ''} ago ({exact_time})"
    elif seconds < 172800:
        return f"yesterday ({exact_time})"
    else:
        days = int(seconds / 86400)
        return f"{days} day{'s' if days != 1 else ''} ago ({exact_time})"