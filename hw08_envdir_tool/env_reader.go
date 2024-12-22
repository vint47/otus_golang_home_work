package main

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var ErrFileName = errors.New("file name invalid")

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	envs := make(Environment)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if strings.ContainsRune(name, '=') {
			return nil, ErrFileName
		}

		content, err := os.ReadFile(filepath.Join(dir, name))
		if err != nil {
			return nil, err
		}
		content = bytes.Replace(content, []byte{0}, []byte{'\n'}, -1)
		value := strings.TrimRight(string(content), " \t\n\r")

		envs[name] = EnvValue{value, len(value) == 0}
	}

	return envs, nil
}
