package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/MagnusTrier/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.args) > 0 {
		n, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return err
		}
		limit = int32(n)
	}

	params := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), params)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s\n", post.PublishedAt.Format("Mon Jan 2"))
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}
