package database

import "github.com/google/uuid"

type Subscription struct {
	ID          uuid.UUID `json:"id"`
	ServiceName string    `json:"service_name"`
	PriceRub    int       `json:"price_rub"`
	UserID      uuid.UUID `json:"user_id"`
	StartDate   string    `json:"start_date"`
	EndDate     string    `json:"end_date"`
}
