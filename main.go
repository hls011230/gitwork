package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Indexhandler(c *gin.Context) {

	c.JSON(http.StatusOK,gin.H{
		"Msg" : "HelloWorld",
	})
}

func Start()  {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/index",Indexhandler)
	r.Run(":8080")
}

func main() {
	Start()
}
