package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Chin-mayyy/Blog_aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *State, cmd Command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.command_name)
	}
	name := cmd.arguments[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func registerHandler(s *State, cmd Command) error {
	if len(cmd.arguments) != 1 {
		return errors.New("Enter a name!!")
	}

	userName := cmd.arguments[0]

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	}

	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	// Add debug print before setting user
	fmt.Println("Before SetUser:", s.cfg.CurrentUserName)

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	// Add debug print after setting user
	fmt.Println("After SetUser:", s.cfg.CurrentUserName)

	return nil

}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}

func GetUsers(s *State, cmd Command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list users: %w", err)
	}

	fmt.Println("current user: ", s.cfg.CurrentUserName)
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %v\n", user.Name)
	}
	return nil

}
