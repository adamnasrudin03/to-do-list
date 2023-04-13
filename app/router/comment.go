package router

import (
	"adamnasrudin03/to-do-list/app/controller"

	"github.com/gin-gonic/gin"
)

func ActivityRouter(e *gin.Engine, h controller.ActivityController) {
	activityRoutes := e.Group("/activity-groups")
	{

		activityRoutes.POST("/", h.Create)
		activityRoutes.GET("/", h.GetAll)
		activityRoutes.PUT("/:id", h.Update)
		activityRoutes.GET("/:id", h.GetOne)
		activityRoutes.DELETE("/:id", h.Delete)
	}
}
