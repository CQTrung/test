package repository

import "gorm.io/gorm"

type Repository struct {
	database *gorm.DB
}
