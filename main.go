package main

import (
	"github.com/Xebec19/scaling-octo-memory/auth"
	"github.com/Xebec19/scaling-octo-memory/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.Connect()

	auth.Routes(r)

	r.Run(":8080")
}
