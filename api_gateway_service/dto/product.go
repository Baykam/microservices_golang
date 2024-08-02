package dto

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Images      *[]string `json:"images"`
	Videos      *[]string `json:"videos"`
	UserID      *string   `json:"user_id"`
}
