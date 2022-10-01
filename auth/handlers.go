package auth

import (
	"net/http"

	"github.com/Xebec19/scaling-octo-memory/util"
	"github.com/gin-gonic/gin"
)

func register(c *gin.Context) {
	var req userRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "done"})
}
