package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, _ command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Encountered error when getting users: %w\n", err)
	}
	for _, u := range users {
		fmt.Printf(" * %v", u.Name)
		if u.Name == s.cfg.CurrentUserName {
			fmt.Printf(" (current)")
		}
		fmt.Print("\n")
	}
	return nil
}
