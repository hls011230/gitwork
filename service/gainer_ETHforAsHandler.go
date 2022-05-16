package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func gainer_ETHforAsHandler(c *gin.Context) {
	gid, _ := strconv.Atoi(c.Request.Header.Get("gid"))
	var EthForAs model.PostETHforAS
	if err := c.ShouldBind(&EthForAs); err != nil {
		serializer.RespError(c, err)
		return
	}

	err := v1.Ganiner_ETHforAs(gid, &EthForAs)
	if err != nil {
		serializer.RespError(c, err)
		return
	}
	serializer.RespOK(c, "兑换成功")
}
