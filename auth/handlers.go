package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{"message": "done"})
}
