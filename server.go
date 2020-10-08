package main


import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()
	r.GET("/ping", Pong)
	r.POST("/SubString", SubString)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}