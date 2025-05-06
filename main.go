package main 

import _ "github.com/lib/pq"

import (
	"os"
	"log"
	"gator/internal/config"
	"gator/internal/database"
	"database/sql"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

func main() {
	config, err := config.Read()	
	
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", config.DBUrl) 
	if err != nil {
		log.Fatal("error getting database conncection")
	}

	dbQueries := database.New(db)

	appState := state{cfg: &config, db: dbQueries} 

	commands := commands{commandList: make(map[string]func(*state, command) error)}

	// Removing first arg which is program name
	args := os.Args

	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		os.Exit(1)
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerGetUsers)
	commands.register("agg", handlerAggregate)
	commands.register("addfeed", handlerAddFeed)
	commands.register("feeds", handlerGetFeeds)
	commands.register("follow", handlerFollowFeeds)
	commands.register("following", handlerGetFeedFollowsForUser)

	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	err = commands.run(&appState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
