from database import db
from datetime import datetime
from flask import jsonify

# First time the agent connects
def check_in(agent_id, agent_username, agent_hostname, agent_os):
    conn = db.get_conn()
    cursor = conn.cursor()

    last_seen = datetime.now()
    
    # Verify if agent_id is already registered
    cursor.execute("""
        SELECT 1 FROM agents WHERE agent_id = ?
    """, (agent_id,))
    exists = cursor.fetchone()

    if not exists:
        cursor.execute("""
            INSERT INTO agents (agent_id, agent_username, agent_hostname, agent_os, last_seen)
            VALUES (?, ?, ?, ?, ?)
        """, (agent_id, agent_username, agent_hostname, agent_os, last_seen))

        conn.commit()
    # If exists, update the lastseen
    else:
        cursor.execute("""
            UPDATE agents
            SET last_seen = ?
            WHERE agent_id = ?
        """, (last_seen, agent_id))
        conn.commit()
    
    conn.close()

# Every x seconds the agent looks for new commands and update the last_seen time
def check_commands(agent_id):
    conn = db.get_conn()
    cursor = conn.cursor()
    
    # search commands
    cursor.execute("""
        SELECT command_id, agent_command, command_datetime
        FROM commands
        WHERE agent_id = ? AND command_result is NULL
    """, (agent_id,))
    commands = cursor.fetchall()
    
    conn.close()
    
    processed_commands = []
    
    for command in commands:
        processed_command = {
            "command_id": command[0],
            "agent_command": command[1],
            "command_datetime": command[2]
        }
        
        processed_commands.append(processed_command)
    
    return jsonify({"commands": processed_commands})

def add_command_result(command_id, command_result):
    conn = db.get_conn()
    cursor = conn.cursor()
    
    cursor.execute("""
        UPDATE commands
        SET command_result = ?
        WHERE command_id = ?
    """, (command_result, command_id))
    conn.commit()
    conn.close()