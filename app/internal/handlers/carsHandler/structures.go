package carsHandler

import (
	"github.com/google/uuid"
	"time"
)

type CarResponse struct {
	ID        uuid.UUID `json:"id"`
	ModelName string    `json:"model_name"`
	Price     int64     `json:"price"`
	Country   string    `json:"country"`
	Producer  string    `json:"producer"`
	Date      time.Time `json:"date"`
}

type CreateCarRequest struct {
	ModelName string    `json:"model_name"`
	Price     int64     `json:"price"`
	Country   string    `json:"country"`
	Producer  string    `json:"producer"`
	Date      time.Time `json:"date"`
}

type UpdateCarRequest struct {
	ModelName string `json:"model_name"`
	Price     int64  `json:"price"`
}
