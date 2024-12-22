package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s /path/to/env/dir command [args...]\n", os.Args[0])
		os.Exit(1)
	}

	dir := os.Args[1]
	command := os.Args[2:]

	envs, err := ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
	}

	RunCmd(command, envs)
}
