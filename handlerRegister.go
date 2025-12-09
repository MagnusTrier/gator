package main

import (
	"context"
	"fmt"
	"github.com/MagnusTrier/gator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("register command requires name\n")
	}
	name := cmd.args[0]

	params := database.CreateUserParams{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ID:        uuid.New(),
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Encountered error when creating user: %w\n", err)
	}

	s.cfg.SetUser(name)

	fmt.Printf("User created | %v\n", user)
	return nil
}
