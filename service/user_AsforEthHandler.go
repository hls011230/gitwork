package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func user_AsforEthHandler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Request.Header.Get("uid"))
	var ASfroETH model.PostETHforAS
	if err := c.ShouldBind(&ASfroETH); err != nil {
		serializer.RespError(c, err)
		return
	}
	err := v1.User_ASforETH(uid, &ASfroETH)
	if err != nil {
		serializer.RespError(c, err)
	}
	serializer.RespOK(c, "兑换成功")
}
