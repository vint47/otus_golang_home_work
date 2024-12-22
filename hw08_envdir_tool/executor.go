package main

import (
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	setupEnv(env)

	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		return 0
	}

	return 1
}

func setupEnv(env Environment) {
	for key, value := range env {
		if value.NeedRemove {
			os.Unsetenv(key)
		} else {
			os.Setenv(key, value.Value)
		}
	}
}
