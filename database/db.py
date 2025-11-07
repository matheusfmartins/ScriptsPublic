import sqlite3

DB_NAME = "database.db"

def get_conn():
    return sqlite3.connect(DB_NAME)

def init_db():
    conn = get_conn()
    cursor = conn.cursor()
    
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS agents (
            agent_id TEXT NOT NULL,
            agent_username TEXT,
            agent_hostname TEXT,
            agent_os TEXT,
            last_seen TEXT
        )
    """)
    
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS commands (
            command_id INTEGER PRIMARY KEY AUTOINCREMENT,
            agent_id TEXT NOT NULL,
            agent_command TEXT,
            command_result TEXT,
            command_datetime TEXT
        )
    """)
    
    cursor.execute("""
        CREATE TABLE IF NOT EXISTS users (
            user_id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL
        )
    """)
    
    print("[INFO] Database created.")
    
    # Verifica se o usuário 'matheus' já existe
    cursor.execute("SELECT * FROM users WHERE username = ?", ("matheus",))
    if cursor.fetchone() is None:
        cursor.execute("INSERT INTO users (username, password) VALUES (?, ?)", (
            "matheus",
            "cfc568108f262ae670b23b6c602caf62ed2fbc290c240fa8f5ff464d89bce0b1"
        ))
        print("[INFO] User matheus created.")
    
    conn.commit()
    conn.close()