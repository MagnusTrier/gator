package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, _ command) error {
	path := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), path)
	if err != nil {
		return err
	}

	fmt.Print(feed)

	return nil
}
