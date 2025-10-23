package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"newsletter-backend/internal/models"
)

type NewsletterRepository interface {
	CreateEmail(email string) error
	EmailExists(email string) (bool, error)
}

type newsletterRepository struct {
	db *sql.DB
}

func NewNewsletterRepository(db *sql.DB) NewsletterRepository {
	return &newsletterRepository{db: db}
}

func (r *newsletterRepository) CreateEmail(email string) error {
	email = strings.ToLower(strings.TrimSpace(email))

	query := `INSERT INTO emails (email) VALUES (?)`

	_, err := r.db.Exec(query, email)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return fmt.Errorf("email already subscribed")
		}
		return fmt.Errorf("failed to insert email: %w", err)
	}

	return nil
}

func (r *newsletterRepository) EmailExists(email string) (bool, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	query := `SELECT COUNT(*) FROM emails WHERE email = ?`

	var count int
	err := r.db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check email existence: %w", err)
	}

	return count > 0, nil
}
