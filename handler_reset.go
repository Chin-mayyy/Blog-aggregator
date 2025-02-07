package main

import (
	"context"
	"fmt"
	"os"
)

func resetHandler(s *State, cmd Command) error {
	if err := s.db.DeleteUser(context.Background()); err != nil {
		os.Exit(1)
		return err
	}
	fmt.Println("Database reset successfully!")
	return nil
}
