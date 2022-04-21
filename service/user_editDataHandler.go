package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/db/model"
	"A11Smile/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

//修改用户名
func user_editUserNameHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"msg": "error",
			"err": err.Error(),
		})
	}

	if err := v1.EditUserName(user.Id, user.Uname); err != nil {
		serializer.RespError(c, err)
		return

	}

	serializer.RespOK(c, "修改成功!")

}

//修改用户简介
func user_editUserResumeHandler(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"msg": "error",
			"err": err.Error(),
		})
	}

	if err := v1.EditUserResume(user.Id, user.Resume); err != nil {
		serializer.RespError(c, err)
		return

	}

	serializer.RespOK(c, "修改成功!")

}

//查询资料
func user_authenticationSeeHandler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Request.Header.Get("uid"))
	user, err := v1.UserAuthenticationSee(uid)

	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, user)
}