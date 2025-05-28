package domain

import "time"

type Character struct {
	ID          string     `json:"id" gorm:"primaryKey;column:id"`
	Affiliation string     `json:"affiliation" gorm:"column:affiliation"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty" gorm:"column:deleted_at"`
	Description string     `json:"description" gorm:"column:description;type:text"`
	Gender      string     `json:"gender" gorm:"column:gender"`
	Image       string     `json:"image" gorm:"column:image"`
	Ki          string     `json:"ki" gorm:"column:ki"`
	MaxKi       string     `json:"maxKi" gorm:"column:max_ki"`
	Name        string     `json:"name" gorm:"column:name"`
	Race        string     `json:"race" gorm:"column:race"`
}
