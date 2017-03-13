// +build !windows

package main

import (
	"os"
	"syscall"
)

func doTheDockers() {
	target := "/Users/caro8994/.dvm/bin/docker/1.12.1/docker"
	env := os.Environ()

	err := syscall.Exec(target, os.Args, env)
	if err != nil {
		panic(err)
	}
}
