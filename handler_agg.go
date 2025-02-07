package main

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %v", err)
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error fetching feed: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading the response: %v", err)
	}

	var feed RSSFeed
	if err := xml.Unmarshal(data, &feed); err != nil {
		return nil, fmt.Errorf("Error unmarshaling the body response: %v", err)
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)

	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}

	return &feed, nil
}

func agg(s *State, cmd Command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("Not enough arguments!")
	}

	time_between_reqs := cmd.arguments[0]

	if time_between_reqs == "0s" {
		return errors.New("Set some time difference")
	}

	complex, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("Error parsing the time: %v", err)
	}

	fmt.Printf("Collecting feeds every %v", complex)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	ticker := time.NewTicker(complex)
	defer ticker.Stop()

	go func() {
		<-sigChan
		fmt.Println("\nReceived Ctrl+C, Exiting gracefully!")
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Shutting down.")
			return nil
		case <-ticker.C:
			fmt.Println("Requesting for the feed...")
			if err := scrapeFeeds(s); err != nil {
				return fmt.Errorf("Error getting the feed: %v", err)
			}
			return nil
		}
	}
}
func scrapeFeeds(s *State) error {
	feedToBeFetched, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Error getting the feed to fetched: %v", err)
	}

	if err := s.db.MarkFeedFetched(context.Background(), feedToBeFetched.ID); err != nil {
		return fmt.Errorf("Error marking the fetched feed: %v", err)
	}

	feed, err := fetchFeed(context.Background(), feedToBeFetched.Url)
	if err != nil {
		return fmt.Errorf("Error fetching the RSSfeed: %v", err)
	}

	fmt.Printf("Channel: %s\n", feed.Channel.Title)
	fmt.Printf("Description: %s\n", feed.Channel.Description)
	fmt.Printf("Link: %s\n\n", feed.Channel.Link)

	// Print items
	for _, item := range feed.Channel.Item {
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Description: %s\n", item.Description)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("Published: %s\n\n", item.PubDate)
	}

	return nil
}
