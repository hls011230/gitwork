package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func gainer_CheckTheAS(c *gin.Context) {

	gid, _ := strconv.Atoi(c.Request.Header.Get("gid"))

	balance, err := v1.Connect6_CheckTheAS(gid)

	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, balance)

}
