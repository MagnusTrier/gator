package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	methods map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	if val, ok := c.methods[cmd.name]; ok {
		if err := val(s, cmd); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("command does not exist\n")
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.methods[name] = f
}
