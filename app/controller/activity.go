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

type ActivityController interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type activityHandler struct {
	Service *service.Services
}

func NewActivityController(srv *service.Services) ActivityController {
	return &activityHandler{
		Service: srv,
	}
}

func (c *activityHandler) Create(ctx *gin.Context) {
	input := dto.CreateActivity{}
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

	activity := entity.Activity{
		Title: input.Title,
		Email: input.Email,
	}

	res, httpStatus, status, err := c.Service.Activity.Create(ctx, activity)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), status, nil))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Success", "Success", res))
}

func (c *activityHandler) GetAll(ctx *gin.Context) {

	res, httpStatus, status, err := c.Service.Activity.GetAll(ctx)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), status, nil))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Success", "Success", res))
}

func (c *activityHandler) GetOne(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), "Bad Request", nil))
		return
	}

	res, httpStatus, status, err := c.Service.Activity.GetByID(ctx, ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), status, nil))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Success", "Success", res))
}

func (c *activityHandler) Update(ctx *gin.Context) {
	input := entity.Activity{}
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

	activity := entity.Activity{
		Title: input.Title,
		Email: input.Email,
	}

	res, httpStatus, status, err := c.Service.Activity.UpdateByID(ctx, ID, activity)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), status, nil))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Success", "Success", res))
}

func (c *activityHandler) Delete(ctx *gin.Context) {
	ID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		err = errors.New("invalid parameter id")
		ctx.JSON(http.StatusBadRequest, helpers.APIResponse(err.Error(), "Bad Request", nil))
		return
	}

	httpStatus, status, err := c.Service.Activity.DeleteByID(ctx, ID)
	if err != nil {
		ctx.JSON(httpStatus, helpers.APIResponse(err.Error(), status, nil))
		return
	}

	ctx.JSON(httpStatus, helpers.APIResponse("Deleted", "Success", nil))
}
