from flask import Flask
from database.db import init_db
from routes.agent_calls import agent_calls
from routes.server_calls import server_calls

app = Flask(__name__)
app.register_blueprint(agent_calls)
app.register_blueprint(server_calls)

if __name__ == "__main__":
    init_db()
    app.run(host="0.0.0.0", port=5000, debug=True)