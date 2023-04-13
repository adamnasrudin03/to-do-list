package app

import (
	"adamnasrudin03/to-do-list/app/repository"

	"gorm.io/gorm"
)

func WiringRepository(db *gorm.DB) *repository.Repositories {
	return &repository.Repositories{
		Activity: repository.NewActivityRepository(db),
	}
}
