package main

import (
	"log"
	"os"
	"time"
	"context"
	"github.com/google/uuid"
	"fmt"
	"gator/internal/database"
)
func handlerAddFeed (s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <name> <feedUrl>", cmd.Name)
	}	

	// Get current user
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)

	if err != nil {
		log.Fatal("Failed to get current user")
		os.Exit(1)
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
		Url: cmd.Args[1],
		UserID: user.ID,
	})

	if err != nil {
		log.Fatal("Failed to create feed")
		os.Exit(1)
	}

	fmt.Println("Feed Successfully Created:")
	fmt.Printf(" * ID:       %v\n", feed.ID)
	fmt.Printf(" * Name:     %v\n", feed.Name)
	fmt.Printf(" * Url:      %v\n", feed.Url)

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	})

	if err != nil {
		log.Fatal("Failed to create feed follow")
		os.Exit(1)
	}

	fmt.Println("Feed Follow Successfully Created:")
	fmt.Printf(" * Feed: %v\n", feedFollow.FeedName)
	fmt.Printf(" * Name: %v\n", feedFollow.UserName)

	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())

	if err != nil {
		log.Fatal("Failed to retrieve feeds")
		os.Exit(1)
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			log.Fatal("Failed to get user")	
			os.Exit(1)
		}

		fmt.Printf("* %v\n", feed.Name)
		fmt.Printf("  %v\n", feed.Url)
		fmt.Printf("  %v\n", user.Name)
	}

	return nil
}

func handlerFollowFeeds(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <feedUrl>", cmd.Name)
	}

	feed, err := s.db.GetFeed(context.Background(), cmd.Args[0])
	if err != nil {
		log.Fatal("Failed to retrieve feed")
		os.Exit(1)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		log.Fatal("Failed to retrieve user")
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

func handlerGetFeedFollowsForUser(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		log.Fatal("Failted to retrieve user")
	}

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
