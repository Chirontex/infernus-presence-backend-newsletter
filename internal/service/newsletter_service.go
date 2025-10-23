package service

import (
	"fmt"
	"strings"

	"newsletter-backend/internal/repository"
	"newsletter-backend/internal/validator"
)

type NewsletterService interface {
	Subscribe(email, clientToken string) error
}

type newsletterService struct {
	repo                repository.NewsletterRepository
	expectedClientToken string
}

func NewNewsletterService(repo repository.NewsletterRepository, clientToken string) NewsletterService {
	return &newsletterService{
		repo:                repo,
		expectedClientToken: clientToken,
	}
}

func (s *newsletterService) Subscribe(email, clientToken string) error {
	// Validate email
	if err := validator.ValidateEmail(email); err != nil {
		return err
	}

	// Validate client token
	if err := validator.ValidateClientToken(clientToken, s.expectedClientToken); err != nil {
		return err
	}

	// Normalize email
	email = strings.ToLower(strings.TrimSpace(email))

	// Check if email already exists
	exists, err := s.repo.EmailExists(email)
	if err != nil {
		return fmt.Errorf("failed to check email: %w", err)
	}

	if exists {
		return fmt.Errorf("email already subscribed")
	}

	// Create email subscription
	if err := s.repo.CreateEmail(email); err != nil {
		return err
	}

	return nil
}
