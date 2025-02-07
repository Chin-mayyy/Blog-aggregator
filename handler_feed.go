package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/Chin-mayyy/Blog_aggregator/internal/database"
)

func addFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.arguments) != 2 {
		os.Exit(1)
		return errors.New("Not enough arguments...")
	}

	params := database.CreateFeedParams{
		Name:   cmd.arguments[0],
		Url:    cmd.arguments[1],
		UserID: user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), params)
	if err != nil {
		return errors.New("couldn't create feed!!!")
	}

	fmt.Println(feed.ID)
	fmt.Println(feed.CreatedAt)
	fmt.Println(feed.UpdatedAt)
	fmt.Println(feed.Name)
	fmt.Println(feed.Url)
	fmt.Println(feed.UserID)

	paramsFollow := database.CreateFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), paramsFollow)
	if err != nil {
		return fmt.Errorf("Error creating follow feed : %v", err)
	}

	return nil
}

func feeds(s *State, cmd Command) error {
	if len(cmd.arguments) != 0 {
		return errors.New("No argument required!")
	}

	feeds, err := s.db.GetFeed(context.Background())
	if err != nil {
		return errors.New("Error getting the feed!")
	}

	for _, feed := range feeds {
		fmt.Printf("Name : %v\n", feed.Name)
		fmt.Printf("Url : %v\n", feed.Url)
		fmt.Printf("Name_2 : %v\n", feed.Name_2)
	}

	return nil
}
