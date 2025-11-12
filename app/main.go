package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func read() (string, error) {
	fmt.Fprint(os.Stdout, "$ ")

	cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return "", err
	}

	return cmd, nil
}

type Builtins struct {
	keywords []string
}

func (u *Builtins) Methods(arg []string) map[string]func() {
	return map[string]func(){
		"exit": func() {
			Exit(arg)
		},
		"type": func() {
			Type(arg, u.keywords)
		},
		"echo": func() {
			Echo(arg)
		},
		"pwd": func() {
			Pwd()
		},
	}
}

func categorizeCmd(cmd string) {
	cmdFields := strings.Fields(cmd)
	command := cmdFields[0]

	builtins := Builtins{keywords: []string{
		"exit", "type", "echo", "pwd",
	}}
	methods := builtins.Methods(cmdFields)

	if slices.Contains(builtins.keywords, cmdFields[0]) {
		methods[cmdFields[0]]()
		fmt.Print("\n")
		return
	}

	if filePath, err := FindFileInBinary(cmdFields[0]); err == nil {
		RunExe(filePath, command, cmdFields...)
		return
	}

	fmt.Printf("%s: command not found\n", strings.TrimSpace(cmd))
}

func main() {
	for {
		cmd, err := read()
		if err != nil {
			os.Exit(0)
		}

		categorizeCmd(cmd)

	}
}
