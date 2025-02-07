package main

import (
	"errors"
)

type Command struct {
	command_name string
	arguments    []string
}

type Commands struct {
	registeredCommand map[string]func(*State, Command) error
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.registeredCommand[name] = f
}

func (c *Commands) run(s *State, cmd Command) error {
	f, ok := c.registeredCommand[cmd.command_name]
	if !ok {
		return errors.New("command not found")
	}

	return f(s, cmd)
}
