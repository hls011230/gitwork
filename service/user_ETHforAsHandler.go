package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func user_ETHforAsHandler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Request.Header.Get("uid"))
	var EthForAs model.PostETHforAS
	if err := c.ShouldBind(&EthForAs); err != nil {
		serializer.RespError(c, err)
		return
	}

	err := v1.User_ETHforAs(uid, &EthForAs)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, "兑换成功")
}
