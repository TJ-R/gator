package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"log"
	"os"
	"strconv"
)

func handlerBrowse(s *state, cmd command, user database.User) error {	
	limit := int32(2)
	if len(cmd.Args) > 0 {
		cmd_limit, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			log.Fatal("Failed to convert limit to int")
			os.Exit(1)
		}

		limit = int32(cmd_limit)
	}

	posts, err := s.db.GetPosts(context.Background(), database.GetPostsParams{
		UserID: user.ID,
		Limit: limit,
	})

	if err != nil {
		log.Fatal("Failed to get posts")
		os.Exit(1)
	}

	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"))
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil
}
