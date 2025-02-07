package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/Chin-mayyy/Blog_aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) (func(s *State, cmd Command) error) {
	return func(s *State, cmd Command) error {
		if s.cfg.CurrentUserName == "" {
			return errors.New("No current user logged in")
		}

		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err !=  nil {
			return fmt.Errorf("Error getting the user : %v", err)
		}

		handler(s, cmd, user)
		return nil
	}
}
