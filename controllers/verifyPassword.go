package controllers

import (
	"net/http"

	"github.com/filipegms5/password-check-restful/models"
	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) {
	var obj models.ObjectRequest

	bindErr := c.ShouldBindJSON(&obj)

	//tratamento de erros
	if bindErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "message": "Request body not valid"})
		return
	}

	verifiedPassword := models.CheckPassword(obj)

	c.JSON(http.StatusOK, gin.H{"result": verifiedPassword})
}
