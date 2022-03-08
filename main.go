package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//err := db.Init()
	//if err!=nil {
	//	panic(err)
	//}
	r := gin.Default()
	r.GET("/index" ,func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{
			"Msg" : "123456" ,
		})
	})
	r.Run(":8080")
	//http.HandleFunc("/", service.IndexHandler)
	//http.HandleFunc("/api/count", service.CounterHandler)
	//
	//log.Fatal(http.ListenAndServe(":80", nil))
}
