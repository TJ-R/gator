package main

import (
	"fmt"
	"log"
	"time"
	"os"
)

func handlerAggregate(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <time_between_reqs> (in format of 1s for 1 second)", cmd.Name)
	}

	timeBetweenReqs, err := time.ParseDuration(cmd.Args[0])

	if err != nil {
		log.Fatal("Failed to parse duration")
		os.Exit(1)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)

	ticker := time.NewTicker(timeBetweenReqs)

	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
