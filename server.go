package main


import "github.com/gin-gonic/gin"


func main() {
	r := setupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	data := _DataHandler{}

	r.GET("/ping", Pong)
	r.POST("/SubString", data.EndPointDataCleansing)

	return r
}