package main

import _ "github.com/lib/pq"

import (
	"database/sql"
	"fmt"
	"github.com/MagnusTrier/gator/internal/config"
	"github.com/MagnusTrier/gator/internal/database"
	"os"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Print(err)
	}

	appState := state{
		cfg: &cfg,
	}

	db, err := sql.Open("postgres", appState.cfg.DBURL)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	dbQueries := database.New(db)
	appState.db = dbQueries

	cmds := commands{
		methods: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

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

	os.Exit(0)
}
