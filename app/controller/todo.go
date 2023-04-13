package controller

import (
	"adamnasrudin03/to-do-list/app/dto"
	"adamnasrudin03/to-do-list/app/entity"
	"adamnasrudin03/to-do-list/app/service"
	"adamnasrudin03/to-do-list/pkg/helpers"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TodoController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type todoHandler struct {
	Service *service.Services
}

func NewTodoController(srv *service.Services) TodoController {
	return &todoHandler{
		Service: srv,
	}
}

func (c *todoHandler) Create(ctx *gin.Context) {
	input := dto.CreateTodo{}
	validate := validator.New()

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), "Bad Request", nil))
		return
	}

	err = validate.Struct(input)
	if err != nil {
		errors := helpers.FormatValidationError(err)

		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(errors, "Bad Request", nil))
		return
	}

	Todo := entity.Todo{
		Title:           input.Title,
		ActivityGroupID: input.ActivityGroupID,
		IsActive:        input.IsActive,
		Priority:        input.Priority,
	}

	res, httpStatus, status, err := c.Service.Todo.Create(ctx, Todo)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), status, nil))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Success", "Success", res))
}

func (c *todoHandler) GetAll(ctx *gin.Context) {
	ActivityGroupID, err := strconv.ParseUint(ctx.Query("activity_group_id"), 10, 32)
	if err != nil {
		err = errors.New("invalid query parmaeter activity_group_id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), "Bad Request", nil))
		return
	}
	if ActivityGroupID == 0 {
		err = errors.New("activity_group_id must be greater than 0")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), "Bad Request", nil))
		return
	}

	res, httpStatus, status, err := c.Service.Todo.GetAllByActivityGroupID(ctx, ActivityGroupID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), status, nil))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Success", "Success", res))
}

func (c *todoHandler) GetOne(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), "Bad Request", nil))
		return
	}

	res, httpStatus, status, err := c.Service.Todo.GetByID(ctx, ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), status, nil))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Success", "Success", res))
}

func (c *todoHandler) Update(ctx *gin.Context) {
	input := dto.UpdateTodo{}
	validate := validator.New()

	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), "Bad Request", nil))
		return
	}

	err = ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), "Bad Request", nil))
		return
	}

	err = validate.Struct(input)
	if err != nil {
		errors := helpers.FormatValidationError(err)

		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(errors, "Bad Request", nil))
		return
	}

	Todo := entity.Todo{
		Title:           input.Title,
		ActivityGroupID: input.ActivityGroupID,
		IsActive:        input.IsActive,
		Priority:        input.Priority,
	}

	res, httpStatus, status, err := c.Service.Todo.UpdateByID(ctx, ID, Todo)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), status, nil))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Success", "Success", res))
}

func (c *todoHandler) Delete(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), "Bad Request", nil))
		return
	}

	httpStatus, status, err := c.Service.Todo.DeleteByID(ctx, ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), status, nil))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Deleted", "Success", nil))
}
