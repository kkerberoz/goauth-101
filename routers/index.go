package routers

import (
	"ginauth101/controllers"
	"ginauth101/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setAuthRoute(router *gin.Engine) {
	authController := new(controllers.AuthController)
	router.POST("/login", authController.Login)
	router.POST("/register", authController.Register)

	authGroup := router.Group("/")
	authGroup.Use(middlewares.Authentication())
	authGroup.GET("/profile", authController.Profile)

}

// InitRoute ..
func InitRoute() *gin.Engine {

	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World!")
	})

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	setAuthRoute(router)
	return router
}
