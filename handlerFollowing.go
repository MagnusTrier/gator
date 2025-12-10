package main

import (
	"context"
	"fmt"

	"github.com/MagnusTrier/gator/internal/database"
)

func handlerFollowing(s *state, _ command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("Encountered error when getting feed follows for user: %w\n", err)
	}

	fmt.Printf("%v's feeds:\n", user.Name)
	for _, f := range feeds {
		fmt.Printf(" - %v\n", f.FeedName)
	}

	return nil
}
