package app

import (
	"adamnasrudin03/to-do-list/app/repository"
	"adamnasrudin03/to-do-list/app/service"

	"gorm.io/gorm"
)

func WiringRepository(db *gorm.DB) *repository.Repositories {
	return &repository.Repositories{
		Activity: repository.NewActivityRepository(db),
		Todo:     repository.NewTodoRepository(db),
	}
}

func WiringService(repo *repository.Repositories) *service.Services {
	return &service.Services{
		Activity: service.NewActivityService(repo.Activity),
		Todo:     service.NewTodoService(repo.Todo, repo.Activity),
	}
}
