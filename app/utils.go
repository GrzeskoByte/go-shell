package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func RunExe(filePath string, command string, args ...string) error {
	cmd := exec.Command(filePath, args[1:]...)

	cmd.Args = args
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func FindFileInBinary(cmd string) (string, error) {
	for _, path := range strings.Split(os.Getenv("PATH"), ":") {
		filePath := path + "/" + cmd
		file, _ := os.Stat(filePath)

		if _, err := os.Stat(filePath); err == nil {

			mode := file.Mode()
			permissions := mode.Perm()
			isExecutable := permissions&0o100 != 0

			if isExecutable {
				return filePath, nil
			}

		}
	}

	return "", errors.New("NO_FILE_FOUND")
}
