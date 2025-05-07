package main

import (
	"context"
	"log"
	"os"
	"fmt"
	"gator/internal/database"
	"time"
	"github.com/google/uuid"
)

func handlerFollowFeeds(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <feedUrl>", cmd.Name)
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		log.Fatal("Failed to retrieve feed")
		os.Exit(1)
	}

	followed_feed, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		log.Fatal("Failed to create new feed follow")
		os.Exit(1)
	}

	fmt.Printf("* Feed: %v\n", followed_feed.FeedName)
	fmt.Printf("* User: %v\n", followed_feed.UserName)

	return nil
}

func handlerGetFeedFollowsForUser(s *state, cmd command, user database.User) error {
	followed_feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		log.Fatal("Failed to retrieve followed feeds")
		os.Exit(1)
	}

	fmt.Println("Followed Feeds:")
	for _, followed_feed := range followed_feeds {
		fmt.Printf("* %v\n", followed_feed.FeedName)
	}

	return nil
}

func handlerRemoveFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <feedUrl>", cmd.Name)
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		log.Fatal("Failed to retrieve feed")
		os.Exit(1)
	}

	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		log.Fatalf("Failed to unfollow %v\n", cmd.Args[0])
		os.Exit(1)
	}

	return nil
}
