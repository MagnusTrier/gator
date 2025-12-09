package main

import (
	"fmt"
	"github.com/MagnusTrier/gator/internal/config"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Print(err)
	}
	appState := state{
		cfg: &cfg,
	}

	cmds := commands{
		methods: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)

	rawArgs := os.Args

	if len(rawArgs) < 2 {
		fmt.Print("no arguments were provided\n")
		os.Exit(1)
	}

	requestedCommand := command{
		name: rawArgs[1],
		args: rawArgs[2:],
	}

	if err := cmds.run(&appState, requestedCommand); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
