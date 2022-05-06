package service

import (
	v1 "A11Smile/api/v1"
	"A11Smile/serializer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func user_showNoTransactionsHandler(c *gin.Context) {
	uid, _ := strconv.Atoi(c.Request.Header.Get("uid"))
	AllTransactionInformation, err := v1.ShowNoTransactionsHandler(uid)
	if err != nil {
		serializer.RespError(c, err)
	}
	serializer.RespOK(c, AllTransactionInformation)
}
