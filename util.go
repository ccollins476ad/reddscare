package main

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// run executes a command with the specified arguments with its stdout and
// stderr set to match the current process. It blocks until the command
// terminates.
func run(cmdName string, args ...string) error {
	cmd := exec.Command(cmdName, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Debugf("%s %v", cmdName, args)
	return cmd.Run()
}
