package api

import (
	"encoding/json"
	"time"
)

type characterResponse struct {
	ID          json.Number `json:"id"`
	Affiliation string      `json:"affiliation"`
	DeletedAt   *time.Time  `json:"deletedAt"`
	Description string      `json:"description"`
	Gender      string      `json:"gender"`
	Image       string      `json:"image"`
	Ki          string      `json:"ki"`
	MaxKi       string      `json:"maxKi"`
	Name        string      `json:"name"`
	Species     string      `json:"race"`
}
