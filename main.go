package main

import (
	"fmt"
	"net/http"

	"adamnasrudin03/to-do-list/app"
	"adamnasrudin03/to-do-list/app/configs"
	"adamnasrudin03/to-do-list/app/controller"
	routers "adamnasrudin03/to-do-list/app/router"
	"adamnasrudin03/to-do-list/pkg/database"
	"adamnasrudin03/to-do-list/pkg/helpers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = database.SetupMySQLConnection()

	repo     = app.WiringRepository(db)
	services = app.WiringService(repo)

	activityController controller.ActivityController = controller.NewActivityController(services)
)

func main() {
	defer database.CloseMySQLConnection(db)
	config := configs.GetInstance()

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, helpers.APIResponse("Welcome my app", "Success", nil))
	})

	// Route here
	routers.ActivityRouter(router, activityController)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, helpers.APIResponse("Page not found", "Not Found", nil))
	})

	listen := fmt.Sprintf(":%v", config.Appconfig.Port)
	router.Run(listen)
}
