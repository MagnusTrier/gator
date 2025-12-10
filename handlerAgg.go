package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/MagnusTrier/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("agg takes 1 argument: time_between_reqs")
	}

	time_between_reqs := cmd.args[0]

	dur, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(dur)

	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)

	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			return err
		}

	}
	return nil
}

func scrapeFeeds(s *state) error {
	next, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	if err := s.db.MarkFeedFetched(context.Background(), next.ID); err != nil {
		return err
	}

	feed, err := fetchFeed(context.Background(), next.Url)
	if err != nil {
		return err
	}

	for _, item := range feed.Channel.Item {
		parseTime, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			return err
		}

		fmt.Print(item.Description)
		param := database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: parseTime,
			FeedID:      next.ID,
		}

		if _, err := s.db.CreatePost(context.Background(), param); err != nil {
			if strings.Contains(fmt.Sprint(err), "duplicate") {
				continue
			} else {
				fmt.Printf("REE: %v", err)
			}
		}

	}

	return nil
}
