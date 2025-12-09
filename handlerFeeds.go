package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, _ command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Encountered error when getting feeds: %w\n", err)
	}

	for _, f := range feeds {

		user, err := s.db.GetUserId(context.Background(), f.UserID)
		if err != nil {
			return fmt.Errorf("Encountered error when getting user: %w\n", err)
		}

		fmt.Printf(" - %v | %v | %v\n", f.Name, f.Url, user.Name)
	}

	return nil
}
