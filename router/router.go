package router

import (
	"github.com/filipegms5/password-check-restful/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/verify", func(c *gin.Context) {
		controllers.Verify(c)
	})

	return router
}
