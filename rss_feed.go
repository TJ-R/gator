package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"gator/internal/database"
	"html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title 		string 		`xml:"title"`
		Link 		string 		`xml:"link"`
		Description string 		`xml:"description"`
		Item 		[]RSSItem 	`xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title 		string `xml:"title"`
	Link 		string `xml:"link"`
	Description string `xml:"description"`
	PubDate 	string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequestWithContext(
		ctx, 
		"GET", 
		feedURL,
		bytes.NewBuffer([]byte("")),
	)
	if err != nil {
		fmt.Errorf("Failed to build request")
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err 
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rssFeed := RSSFeed{} 
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return nil, err 
	}
	
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssFeed.Channel.Item[i] = item
	}

	return &rssFeed, nil
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Fatal("Failed to get next feed to fetch")
		os.Exit(1)
	}

	s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
		UpdatedAt: time.Now(),
	})

	fmt.Printf("Fetching %v...\n", feed.Url)

	fetchedFeed, err := fetchFeed(context.Background(), feed.Url)

	if err != nil {
		return err
	}

	for _, item := range fetchedFeed.Channel.Item {
		publishedAt := sql.NullTime{}
		if pubTime, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time: pubTime,
				Valid: true,
			}
		}

		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: item.Title,
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid: true,
			},
			PublishedAt: publishedAt,
			FeedID: feed.ID,
		})

		if err != nil && !strings.Contains(err.Error(), "duplicate") {
			fmt.Printf("%v\n", err.Error())
		} 

		if err == nil {
			fmt.Printf("Post for %v created\n", post.Title)
		}
	}

	return nil
}
