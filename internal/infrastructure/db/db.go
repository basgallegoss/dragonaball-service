package db

import (
	"errors"

	"github.com/basgallegoss/dragonball-service/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(dsn string) (*PostgresRepository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&domain.Character{}); err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) Save(c domain.Character) error {
	return r.db.Create(&c).Error
}

func (r *PostgresRepository) FindByName(name string) (*domain.Character, error) {
	var c domain.Character
	err := r.db.Where("name = ?", name).First(&c).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &c, nil

}
