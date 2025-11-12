package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func Echo(cmdFields []string) {
	for _, element := range cmdFields[1:] {
		fmt.Print(element + " ")
	}
}

func Pwd() {
	path, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Print(path)
}

func Type(cmdFields []string, builtins []string) {
	if slices.Contains(builtins, cmdFields[1]) {
		fmt.Printf("%s is a shell builtin", cmdFields[1])
		return
	}

	if path, err := FindFileInBinary(cmdFields[1]); err == nil {
		fmt.Fprintf(os.Stdout, "%s is %s", cmdFields[1], path)
		return
	}

	fmt.Printf("%s: not found", strings.TrimSpace(cmdFields[1]))
}

func Exit(cmdFields []string) {
	if cmdFields[1] == "0" {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
