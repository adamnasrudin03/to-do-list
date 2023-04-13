package service

import (
	"adamnasrudin03/to-do-list/app/entity"
	"adamnasrudin03/to-do-list/app/repository"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ActivityService interface {
	Create(ctx *gin.Context, input entity.Activity) (res entity.Activity, statusCode int, status string, err error)
	GetAll(ctx *gin.Context) (result []entity.Activity, statusCode int, status string, err error)
	GetByID(ctx *gin.Context, ID uint64) (result entity.Activity, statusCode int, status string, err error)
	UpdateByID(ctx *gin.Context, ID uint64, input entity.Activity) (result entity.Activity, statusCode int, status string, err error)
	DeleteByID(ctx *gin.Context, ID uint64) (statusCode int, status string, err error)
}

type activitySrv struct {
	ActivityRepository repository.ActivityRepository
}

func NewActivityService(ActivityRepo repository.ActivityRepository) ActivityService {
	return &activitySrv{
		ActivityRepository: ActivityRepo,
	}
}

func (srv *activitySrv) Create(ctx *gin.Context, input entity.Activity) (res entity.Activity, statusCode int, status string, err error) {
	res, err = srv.ActivityRepository.Create(ctx, input)
	if err != nil {
		log.Printf("[ActivityService-Create] error create data: %+v \n", err)
		return res, http.StatusInternalServerError, "Internal Server Error", err
	}

	return res, http.StatusCreated, "Success", nil
}

func (srv *activitySrv) GetAll(ctx *gin.Context) (result []entity.Activity, statusCode int, status string, err error) {
	result, err = srv.ActivityRepository.GetAll(ctx)
	if err != nil {
		log.Printf("[ActivityService-GetAll] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, "Internal Server Error", err
	}

	return result, http.StatusOK, "Success", nil
}

func (srv *activitySrv) GetByID(ctx *gin.Context, ID uint64) (result entity.Activity, statusCode int, status string, err error) {
	result, err = srv.ActivityRepository.GetByID(ctx, ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || result.ID == 0 {
		message := fmt.Sprintf("Activity with ID %v Not Found", ID)
		err = errors.New(message)
		return result, http.StatusNotFound, "Not Found", err
	}

	if err != nil {
		log.Printf("[ActivityService-GetByID] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, "Internal Server Error", err
	}

	return result, http.StatusOK, "Success", nil
}

func (srv *activitySrv) UpdateByID(ctx *gin.Context, ID uint64, input entity.Activity) (result entity.Activity, statusCode int, status string, err error) {
	sm, err := srv.ActivityRepository.GetByID(ctx, ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || sm.ID == 0 {
		message := fmt.Sprintf("Activity with ID %v Not Found", ID)
		err = errors.New(message)
		return result, http.StatusNotFound, "Not Found", err
	}

	if err != nil {
		log.Printf("[ActivityService-UpdateByID] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, "Internal Server Error", err
	}

	result, err = srv.ActivityRepository.UpdateByID(ctx, ID, input)
	if err != nil {
		log.Printf("[ActivityService-UpdateByID] error update data repo: %+v \n", err)
		return result, http.StatusInternalServerError, "Internal Server Error", err
	}

	return result, http.StatusOK, "Success", nil
}

func (srv *activitySrv) DeleteByID(ctx *gin.Context, ID uint64) (statusCode int, status string, err error) {
	sm, err := srv.ActivityRepository.GetByID(ctx, ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || sm.ID == 0 {
		message := fmt.Sprintf("Activity with ID %v Not Found", ID)
		err = errors.New(message)
		return http.StatusNotFound, "Not Found", err
	}

	if err != nil {
		log.Printf("[ActivityService-DeleteByID] error get data repo: %+v \n", err)
		return http.StatusInternalServerError, "Internal Server Error", err
	}

	err = srv.ActivityRepository.DeleteByID(ctx, ID)
	if err != nil {
		log.Printf("[ActivityService-DeleteByID] error delete data repo: %+v \n", err)
		return http.StatusInternalServerError, "Internal Server Error", err
	}

	return http.StatusOK, "Success", nil
}
