package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("Encountered error when deleting users: %w\n", err)
	}
	fmt.Printf("users deleted successfully\n")
	return nil
}
