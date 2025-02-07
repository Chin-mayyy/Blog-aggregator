package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/Chin-mayyy/Blog_aggregator/internal/database"
)

func follow(s *State, cmd Command) error {
	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("No user logged in currently!")
	}

	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Not enough arguments!")
	}

	url := cmd.arguments[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error getting new feed! : %v", err)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error getting the current user! : %v", err)
	}

	params := database.CreateFeedFollowParams{
		UserID: currentUser.ID,
		FeedID: feed.ID,
	}

	feedFollowRecord, err := s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("Error creating the follow feed record! :%v", err)
	}

	fmt.Println(feedFollowRecord.FeedName)
	fmt.Println(feedFollowRecord.UserName)

	return nil
}

func following(s *State, cmd Command) error {
	if len(cmd.arguments) != 0 {
		return errors.New("Not enough arguments")
	}

	if s.cfg.CurrentUserName == "" {
		return errors.New("No user logged in")
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error getting the current user : %v", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), currentUser.ID)
	if err != nil {
		return fmt.Errorf("Error getting the follow feed : %v", err)
	}

	for _, feed := range feedFollows {
		fmt.Printf("%v\n", feed.FeedName)
	}

	return nil
}

func unfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return errors.New("Not enough arguments!")
	}

	url := cmd.arguments[0]

	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return errors.New("Error getting the feed!")
	}

	params := database.DeleteFeedRecordParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	if err := s.db.DeleteFeedRecord(context.Background(), params); err != nil {
		return fmt.Errorf("Error deleting the record : %w", err)
	}

	return nil
}
