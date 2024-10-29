package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

type shell struct {
	commands [][]string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Starting UNIX shell")

	for {
		root, _ := os.Getwd()
		fmt.Print("~" + root + "> ")
		buffer, _ := reader.ReadString('\n')

		cmds := strings.Split(buffer, "|")

		sh := &shell{}
		trimmed := trimdata(cmds)
		sh.commands = trim2d(trimmed)
		for _, command := range sh.commands {
			switch command[0] {
			case "quit":
				fmt.Println("quit from shell...")
				os.Exit(1)
			case "echo":
				echo(command)
			case "pwd":
				if len(command) > 1 {
					fmt.Println("pwd cannot have more than one argument")
					continue
				}
				fmt.Println(root)
			case "cd":
				if len(command) < 2 {
					err := os.Chdir("/Users")
					if err != nil {
						continue
					}
					continue
				}
				err := os.Chdir(command[1])
				if err != nil {
					fmt.Printf("cd: no such file or directory: %s\n", command[1])
				}
			case "ps":
				process()
			case "kill":
				if len(command) < 2 {
					fmt.Println("kill: not enough arguments")
				}
				pid, err := strconv.Atoi(command[1])
				if err != nil {
					fmt.Printf("cannot convert %s to pid\n", command[1])
					continue
				}
				err = kill(pid)
				if err != nil {
					fmt.Printf("cannot find pid: %v\n", pid)
					continue
				}
				fmt.Printf("killed %v\n", pid)
			default:
				continue
			}
		}
	}

}

func echo(command []string) {
	if len(command) == 1 {
		fmt.Println()
		return
	}
	fmt.Println(strings.Join(command[1:], " "))
}

func trimdata(data []string) []string {
	commands := make([]string, 0)
	for _, cmd := range data {
		cmd = strings.TrimSuffix(cmd, "\n")
		cmd = strings.TrimSuffix(cmd, "\r")
		cmd = strings.TrimSuffix(cmd, " ")
		cmd = strings.TrimPrefix(cmd, " ")
		cmd = strings.TrimSpace(cmd)
		commands = append(commands, cmd)
	}
	return commands
}

func trim2d(data []string) [][]string {
	commands := [][]string{}

	for _, command := range data {
		splitted := strings.Split(command, " ")
		cmd := trimdata(splitted)
		commands = append(commands, cmd)
	}
	return commands
}

func process() {
	processList, err := ps.Processes()
	if err != nil {
		fmt.Println("ps.Processes() failed, are you using windows?")
		return
	}
	var process ps.Process
	for p := range processList {
		process = processList[p]
		fmt.Printf("%d\t%s\n", process.Pid(), process.Executable())
	}
}

func kill(pid int) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	process.Kill()
	return nil
}
