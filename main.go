package main

import (
	"github.com/alfannurfaiz9/koda-b9-gin.git/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.MainRouter(r)

	r.Run("localhost:9000")
}
