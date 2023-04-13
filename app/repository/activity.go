package repository

import (
	"log"

	"adamnasrudin03/to-do-list/app/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ActivityRepository interface {
	Create(ctx *gin.Context, input entity.Activity) (res entity.Activity, err error)
	GetAll(ctx *gin.Context) (result []entity.Activity, err error)
	GetByID(ctx *gin.Context, ID uint64) (result entity.Activity, err error)
	UpdateByID(ctx *gin.Context, ID uint64, input entity.Activity) (result entity.Activity, err error)
	DeleteByID(ctx *gin.Context, ID uint64) (err error)
}

type activityRepo struct {
	DB *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepo{
		DB: db,
	}
}

func (repo *activityRepo) Create(ctx *gin.Context, input entity.Activity) (res entity.Activity, err error) {
	query := repo.DB.WithContext(ctx)

	if err = query.Create(&input).Error; err != nil {
		log.Printf("[ActivityRepository-Create] error Create new data: %+v \n", err)
		return input, err
	}

	return input, err
}

func (repo *activityRepo) GetAll(ctx *gin.Context) (result []entity.Activity, err error) {
	query := repo.DB.WithContext(ctx)

	err = query.Find(&result).Error
	if err != nil {
		log.Printf("[ActivityRepository-GetAll] error get data: %+v \n", err)
		return
	}

	return
}

func (repo *activityRepo) GetByID(ctx *gin.Context, ID uint64) (result entity.Activity, err error) {
	query := repo.DB.WithContext(ctx)

	err = query.Where("id = ?", ID).Take(&result).Error
	if err != nil {
		log.Printf("[ActivityRepository-GetByID][%v] error: %+v \n", ID, err)
		return result, err
	}

	return result, err
}

func (repo *activityRepo) UpdateByID(ctx *gin.Context, ID uint64, input entity.Activity) (result entity.Activity, err error) {
	query := repo.DB.WithContext(ctx)

	err = query.Clauses(clause.Returning{}).Model(&result).Where("id = ?", ID).Updates(entity.Activity{Title: input.Title, Email: input.Email}).Error
	if err != nil {
		log.Printf("[ActivityRepository-UpdateByID][%v] error: %+v \n", ID, err)
		return result, err
	}

	if result.ID == 0 {
		err = query.Where("id = ?", ID).Take(&result).Error
		if err != nil {
			log.Printf("[ActivityRepository-UpdateByID][%v] returning data: %+v \n", ID, err)
			return result, err
		}
	}

	return result, err
}

func (repo *activityRepo) DeleteByID(ctx *gin.Context, ID uint64) (err error) {
	query := repo.DB.WithContext(ctx)

	if err = query.Where("id = ?", ID).Delete(&entity.Activity{}).Error; err != nil {
		log.Printf("[ActivityRepository-DeleteByID][%v] error: %+v \n", ID, err)
		return
	}

	return
}
