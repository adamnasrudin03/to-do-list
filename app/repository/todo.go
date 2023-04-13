package repository

import (
	"log"

	"adamnasrudin03/to-do-list/app/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TodoRepository interface {
	Create(ctx *gin.Context, input entity.Todo) (res entity.Todo, err error)
	GetAllByActivityGroupID(ctx *gin.Context, activityGroupID uint64) (result []entity.Todo, err error)
	GetByID(ctx *gin.Context, ID uint64) (result entity.Todo, err error)
	UpdateByID(ctx *gin.Context, ID uint64, input entity.Todo) (result entity.Todo, err error)
	DeleteByID(ctx *gin.Context, ID uint64) (err error)
}

type todoRepo struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepo{
		DB: db,
	}
}

func (repo *todoRepo) Create(ctx *gin.Context, input entity.Todo) (res entity.Todo, err error) {
	query := repo.DB.WithContext(ctx)

	if err = query.Create(&input).Error; err != nil {
		log.Printf("[TodoRepository-Create] error Create new data: %+v \n", err)
		return input, err
	}

	return input, err
}

func (repo *todoRepo) GetAllByActivityGroupID(ctx *gin.Context, activityGroupID uint64) (result []entity.Todo, err error) {
	query := repo.DB.WithContext(ctx)

	err = query.Where("activity_group_id = ?", activityGroupID).Find(&result).Error
	if err != nil {
		log.Printf("[TodoRepository-GetAllByActivityGroupID] error get data: %+v \n", err)
		return
	}

	return
}

func (repo *todoRepo) GetByID(ctx *gin.Context, ID uint64) (result entity.Todo, err error) {
	query := repo.DB.WithContext(ctx)

	err = query.Where("id = ?", ID).Take(&result).Error
	if err != nil {
		log.Printf("[TodoRepository-GetByID][%v] error: %+v \n", ID, err)
		return result, err
	}

	return result, err
}

func (repo *todoRepo) UpdateByID(ctx *gin.Context, ID uint64, input entity.Todo) (result entity.Todo, err error) {
	query := repo.DB.WithContext(ctx)

	err = query.Clauses(clause.Returning{}).Model(&result).Where("id = ?", ID).
		Updates(entity.Todo{Title: input.Title, IsActive: input.IsActive, Priority: input.Priority}).Error
	if err != nil {
		log.Printf("[TodoRepository-UpdateByID][%v] error: %+v \n", ID, err)
		return result, err
	}

	if result.ID == 0 {
		err = query.Where("id = ?", ID).Take(&result).Error
		if err != nil {
			log.Printf("[TodoRepository-UpdateByID][%v] returning data: %+v \n", ID, err)
			return result, err
		}
	}

	return result, err
}

func (repo *todoRepo) DeleteByID(ctx *gin.Context, ID uint64) (err error) {
	query := repo.DB.WithContext(ctx)

	if err = query.Where("id = ?", ID).Delete(&entity.Todo{}).Error; err != nil {
		log.Printf("[TodoRepository-DeleteByID][%v] error: %+v \n", ID, err)
		return
	}

	return
}
