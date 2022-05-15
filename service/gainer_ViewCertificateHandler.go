package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func gainer_ViewCertificateHandler(c *gin.Context) {
	gid, _ := strconv.Atoi(c.Request.Header.Get("gid"))
	medicalName := c.Query("medicalName")
	usercertificate, err := v1.Gainer_ViewCertificate(gid, medicalName)
	if err != nil {
		serializer.RespError(c, err)
	}
	serializer.RespOK(c, usercertificate)

}
