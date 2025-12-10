package main

import (
	"context"
	"fmt"

	"github.com/MagnusTrier/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("unfollow command takes one command: url\n")
	}

	url := cmd.args[0]

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	if err := s.db.DeleteFeedFollow(context.Background(), params); err != nil {
		return err
	}

	fmt.Printf("successfully unfollowed feed: %v\n", feed.Name)

	return nil
}
