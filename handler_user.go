package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"
	"github.com/google/uuid"
	"log"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), cmd.Args[0])	
	
	if err != nil {
		log.Fatalf("%s does not exist", cmd.Args[0])
		os.Exit(1)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return err
	}

	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name: cmd.Args[0],
	})

	if err != nil {
		log.Fatal("Failed to create user")
		os.Exit(1)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("Failed to set newly create user %s as current user in config", user.Name)
	}

	fmt.Println("User Successfully Created:")
	fmt.Printf(" * ID:       %v\n", user.ID)
	fmt.Printf(" * NAME:     %v\n", user.Name)

	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())

	if err != nil {
		log.Fatal("Failed to retrieve users")
		os.Exit(1)
	}

	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
