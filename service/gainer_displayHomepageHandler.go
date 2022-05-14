package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func gainer_displayHomepageHandler(c *gin.Context) {
	gid, _ := strconv.Atoi(c.Request.Header.Get("gid"))
	gainermsg, err := v1.DisplayGainerHomepage(gid)
	if err != nil {
		serializer.RespError(c, err)
	}
	serializer.RespOK(c, gainermsg)

}
