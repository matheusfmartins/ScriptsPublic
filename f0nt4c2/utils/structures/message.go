package utils

type Message struct {
	AgentId       	string
	AgentHostname 	string
	AgentPWD      	string
	LastConnTime	int
	AgentActive		bool
	AgentOS			string
	Commands      	[]Command
}
