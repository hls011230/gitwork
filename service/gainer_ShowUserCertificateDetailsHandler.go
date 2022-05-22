package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"

	"github.com/gin-gonic/gin"
)

func gainer_ShowUserCertificateDetailsHandler(c *gin.Context) {
	serial := c.Query("serial")
	res, err := v1.ShowUserCertificateDetails(serial)
	if err != nil {
		serializer.RespError(c, err)
		return
	}

	serializer.RespOK(c, res)

}
