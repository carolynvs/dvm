// +build windows

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func doTheDockers() {
	target := `C:\Users\caro8994\.dvm\bin\docker\1.11.2\docker`
	args := os.Args[1:len(os.Args)]

	cmd := exec.Command(target, args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	defer stdin.Close()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	err = cmd.Wait()
	if err != nil {
		if extErr, ok := err.(*exec.ExitError); ok {
			if status, ok := extErr.Sys().(syscall.WaitStatus); ok {
				os.Exit(status.ExitStatus())
			}
		}
		panic(err)
	} else {

		status := cmd.ProcessState.Sys().(syscall.WaitStatus)
		os.Exit(status.ExitStatus())
	}
}
