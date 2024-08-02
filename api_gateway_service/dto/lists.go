package dto

import uuid "github.com/satori/go.uuid"

type product struct {
	ID     uuid.UUID  `json:"id"`
	Title  string     `json:"title"`
	UserID *uuid.UUID `json:"user_id"`
}

type ProductList struct {
	Products []product `json:"products"`
}
