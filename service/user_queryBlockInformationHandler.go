package service

import (
	"A11Smile/eth"
	"A11Smile/serializer"
	"context"
	"math/big"

	"github.com/gin-gonic/gin"
)

func user_queryBlockInformationHandler(c *gin.Context) {

	num, _ := eth.Client.BlockNumber(context.Background())

	var res []interface{}
	for i := 9200; i < int(num); i++ {
		e, _ := eth.Client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if e.Transactions().Len() != 0 {
			for _, v := range e.Transactions() {
				r := struct {
					Value    interface{}
					To       interface{}
					GasPrice interface{}
					Hash     interface{}
					BlockNum interface{}
					Nonce    interface{}
				}{
					Value:    v.Value(),
					To:       v.To(),
					GasPrice: v.GasPrice(),
					Hash:     v.Hash(),
					BlockNum: i,
					Nonce:    v.Nonce(),
				}

				res = append(res, r)
			}
		}
	}

	serializer.RespOK(c, res)

}
