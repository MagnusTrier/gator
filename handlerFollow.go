package main

import (
	"context"
	"fmt"
	"time"

	"github.com/MagnusTrier/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("follow command must have url")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Encountered error when getting feed by url: %w\n", err)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Encountered error when getting user: %w\n", err)
	}

	param := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), param)
	if err != nil {
		return fmt.Errorf("Encountered error when creating feed follow: %w\n", err)
	}

	fmt.Printf("%v --> %v\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}
