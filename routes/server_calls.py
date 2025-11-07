from flask import Blueprint, request, jsonify, abort
from model import server_calls_model, user_model

api_key = "f3a6d88bdfc14c98a2a0f754b4dc9ecda63f31cbe91e9f1a4e7923b08e3b6c61"

server_calls = Blueprint("server_calls", __name__)

def require_api_key(f):
    from functools import wraps
    @wraps(f)
    def decorated(*args, **kwargs):
        key = request.headers.get('X-API-Key')
        if not key or key != api_key:
            abort(401, description="Nothing here.")
        return f(*args, **kwargs)
    return decorated

@server_calls.route("/get_agents", methods=["GET"])
@require_api_key
def get_agents():
    response = server_calls_model.get_agents()
    return response, 201

@server_calls.route("/get_agent_by_id", methods=["POST"])
@require_api_key
def get_agent_by_id():
    data = request.get_json()
    response = server_calls_model.get_agent_by_id(data["agent_id"])
    return response, 201

@server_calls.route("/get_commands", methods=["POST"])
@require_api_key
def get_commands():
    data = request.get_json()
    response = server_calls_model.get_commands(data["agent_id"])
    return response, 201

@server_calls.route("/create_command", methods=["POST"])
@require_api_key
def create_command():
    data = request.get_json()
    response = server_calls_model.create_command(data["agent_id"], data["agent_command"])
    return response, 201

@server_calls.route('/login', methods=['POST'])
def login():
    data = request.get_json()
    username = data['username']
    password = data['password']
    
    user = user_model.authenticate(username, password)
    return user, 201

@server_calls.route("/delete_inactive_agents", methods=["GET"])
@require_api_key
def delete_inactive_agents():
    response = server_calls_model.delete_inactive_agents()
    return response
    