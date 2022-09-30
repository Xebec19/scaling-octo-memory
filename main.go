package main

import (
	"github.com/Xebec19/scaling-octo-memory/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	auth.Routes(r)

	r.Run(":8080")
}
