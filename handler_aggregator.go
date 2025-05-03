package main

import (
	"context"
	"fmt"
	"log"
)

func handlerAggregate (s *state, cmd command) error {
	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		log.Fatalf("error fetching feed: %v", err)
	}

	fmt.Printf("Feed: %+v\n", rssFeed)

	return nil
}


