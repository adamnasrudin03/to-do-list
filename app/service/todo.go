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

type TodoService interface {
	Create(ctx *gin.Context, input entity.Todo) (res entity.Todo, statusCode int, status string, err error)
	GetAllByActivityGroupID(ctx *gin.Context, activityGroupID uint64) (result []entity.Todo, statusCode int, status string, err error)
	GetByID(ctx *gin.Context, ID uint64) (result entity.Todo, statusCode int, status string, err error)
	UpdateByID(ctx *gin.Context, ID uint64, input entity.Todo) (result entity.Todo, statusCode int, status string, err error)
	DeleteByID(ctx *gin.Context, ID uint64) (statusCode int, status string, err error)
}

type todoSrv struct {
	TodoRepository     repository.TodoRepository
	ActivityRepository repository.ActivityRepository
}

func NewTodoService(todoRepo repository.TodoRepository, activityRepo repository.ActivityRepository) TodoService {
	return &todoSrv{
		ActivityRepository: activityRepo,
		TodoRepository:     todoRepo,
	}
}

func (srv *todoSrv) Create(ctx *gin.Context, input entity.Todo) (res entity.Todo, statusCode int, status string, err error) {
	temp, err := srv.ActivityRepository.GetByID(ctx, input.ActivityGroupID)
	if errors.Is(err, gorm.ErrRecordNotFound) || temp.ID == 0 {
		message := fmt.Sprintf("Activity group with ID %v Not Found", input.ActivityGroupID)
		err = errors.New(message)
		return res, http.StatusNotFound, "Not Found", err
	}

	if err != nil {
		log.Printf("[TodoService-Create] error get data repo activity: %+v \n", err)
		return res, http.StatusInternalServerError, "Internal Server Error", err
	}

	res, err = srv.TodoRepository.Create(ctx, input)
	if err != nil {
		log.Printf("[TodoService-Create] error create data: %+v \n", err)
		return res, http.StatusInternalServerError, "Internal Server Error", err
	}

	return res, http.StatusCreated, "Success", nil
}

func (srv *todoSrv) GetAllByActivityGroupID(ctx *gin.Context, activityGroupID uint64) (result []entity.Todo, statusCode int, status string, err error) {
	result, err = srv.TodoRepository.GetAllByActivityGroupID(ctx, activityGroupID)
	if err != nil {
		log.Printf("[TodoService-GetAllByActivityGroupID] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, "Internal Server Error", err
	}

	return result, http.StatusOK, "Success", nil
}

func (srv *todoSrv) GetByID(ctx *gin.Context, ID uint64) (result entity.Todo, statusCode int, status string, err error) {
	result, err = srv.TodoRepository.GetByID(ctx, ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || result.ID == 0 {
		message := fmt.Sprintf("Todo with ID %v Not Found", ID)
		err = errors.New(message)
		return result, http.StatusNotFound, "Not Found", err
	}

	if err != nil {
		log.Printf("[TodoService-GetByID] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, "Internal Server Error", err
	}

	return result, http.StatusOK, "Success", nil
}

func (srv *todoSrv) UpdateByID(ctx *gin.Context, ID uint64, input entity.Todo) (result entity.Todo, statusCode int, status string, err error) {
	temp, err := srv.TodoRepository.GetByID(ctx, ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || temp.ID == 0 {
		message := fmt.Sprintf("Todo with ID %v Not Found", ID)
		err = errors.New(message)
		return result, http.StatusNotFound, "Not Found", err
	}

	if err != nil {
		log.Printf("[TodoService-UpdateByID] error get data repo: %+v \n", err)
		return result, http.StatusInternalServerError, "Internal Server Error", err
	}

	result, err = srv.TodoRepository.UpdateByID(ctx, ID, input)
	if err != nil {
		log.Printf("[TodoService-UpdateByID] error update data repo: %+v \n", err)
		return result, http.StatusInternalServerError, "Internal Server Error", err
	}

	return result, http.StatusOK, "Success", nil
}

func (srv *todoSrv) DeleteByID(ctx *gin.Context, ID uint64) (statusCode int, status string, err error) {
	temp, err := srv.TodoRepository.GetByID(ctx, ID)
	if errors.Is(err, gorm.ErrRecordNotFound) || temp.ID == 0 {
		message := fmt.Sprintf("Todo with ID %v Not Found", ID)
		err = errors.New(message)
		return http.StatusNotFound, "Not Found", err
	}

	if err != nil {
		log.Printf("[TodoService-DeleteByID] error get data repo: %+v \n", err)
		return http.StatusInternalServerError, "Internal Server Error", err
	}

	err = srv.TodoRepository.DeleteByID(ctx, ID)
	if err != nil {
		log.Printf("[TodoService-DeleteByID] error delete data repo: %+v \n", err)
		return http.StatusInternalServerError, "Internal Server Error", err
	}

	return http.StatusOK, "Success", nil
}
