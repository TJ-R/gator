package main

import (
	"errors"
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	commandList map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	cmdFunc, ok := c.commandList[cmd.Name]
	if !ok {
		return errors.New(fmt.Sprintf("%s is not a command.", cmd.Args[0]))
	}

	return cmdFunc(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandList[name] = f
	return
}
