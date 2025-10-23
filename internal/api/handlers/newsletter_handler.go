package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"newsletter-backend/internal/models"
	"newsletter-backend/internal/service"
	"newsletter-backend/internal/validator"
)

type NewsletterHandler struct {
	service service.NewsletterService
}

func NewNewsletterHandler(service service.NewsletterService) *NewsletterHandler {
	return &NewsletterHandler{service: service}
}

func (h *NewsletterHandler) Subscribe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.SubscribeRequest

	// Decode request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode request: %v", err)
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Subscribe
	if err := h.service.Subscribe(req.Email, req.ClientToken); err != nil {
		// Check if it's a validation error
		if _, ok := err.(validator.ValidationError); ok {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Check if it's a business logic error (e.g., duplicate email)
		errMsg := err.Error()
		if errMsg == "email already subscribed" || errMsg == "invalid clientToken" {
			respondWithError(w, http.StatusBadRequest, errMsg)
			return
		}

		// Internal server error
		log.Printf("Internal error during subscription: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	// Success response
	respondWithSuccess(w, http.StatusOK, "Successfully subscribed to newsletter")
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(models.Response{
		Success: false,
		Error:   message,
	})
}

func respondWithSuccess(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(models.Response{
		Success: true,
		Message: message,
	})
}
