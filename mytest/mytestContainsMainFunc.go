package main

import (
	"awesomeProject/mypkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	mypkg.MyFunc()
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello %s", "1")
	})
	router.Run()
}
