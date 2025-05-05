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
