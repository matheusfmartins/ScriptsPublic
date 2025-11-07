from database import db
from flask import jsonify
import hashlib

def authenticate(username, password):
    conn = db.get_conn()
    cursor = conn.cursor()
    
    hashed_password = hash_password(password)
    
    cursor.execute("SELECT user_id, username FROM users WHERE username = ? AND password = ?", (username, hashed_password))
    user_query = cursor.fetchone()
    conn.close()
    
    user = {}
    
    if user_query is not None:
        user = {
            "user_id": user_query[0],
            "username": user_query[1]
        }
    
    return jsonify(user)

def hash_password(password):
    return hashlib.sha256(password.encode()).hexdigest()