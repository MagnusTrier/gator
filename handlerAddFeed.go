package main

import (
	"context"
	"fmt"
	"github.com/MagnusTrier/gator/internal/database"
	"github.com/google/uuid"
	"time"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("addfeed command requires feed name and feed url\n")
	}

	name := cmd.args[0]
	url := cmd.args[1]

	args := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), args)
	if err != nil {
		return fmt.Errorf("Encountered error when creating feed: %w\n", err)
	}

	param := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	if _, err := s.db.CreateFeedFollow(context.Background(), param); err != nil {
		return fmt.Errorf("Encountered error when creating feed follow for user: %w\n", err)
	}

	fmt.Print(feed)
	return nil
}
