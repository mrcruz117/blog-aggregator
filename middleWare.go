package main

import (
	"context"
	"fmt"

	"github.com/mrcruz117/blog-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		// Check if a user is logged in
		currentUserName := s.cfg.CurrentUserName
		if currentUserName == "" {
			// If no user is logged in, return an error
			return fmt.Errorf("no user is logged in")
		}

		// Fetch the user from the database
		user, err := s.db.GetUser(context.Background(), currentUserName)
		if err != nil {
			return fmt.Errorf("failed to fetch the logged-in user: %w", err)
		}

		// Call the wrapped handler, passing the logged-in user
		return handler(s, cmd, user)
	}
}
