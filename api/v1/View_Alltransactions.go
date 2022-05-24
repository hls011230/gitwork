package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func ShowAllTransactionsHandler(id int) ([]interface{}, error) {
	// 根据uid获取用户的私钥和地址
	DB := db.Get()
	var w model.Wallet
	DB.Table("users").First(&w, "id = ?", id)
	r1, r2, err := eth.Ins.UserNSeeCertificateState(&bind.CallOpts{Context: context.Background(), From: common.HexToAddress(w.BlockAddress)})

	var r []interface{}
	for k, v := range r1 {
		if v.HospitalName != "" {
			var g model.Gainer
			DB.Table("gainers").First(&g, "block_address = ?", v.Soliciter.String())
			r1 := struct {
				User         common.Address
				Soliciter    common.Address
				HospitalName string
				MedicalName  string
				Certificate  string
				Amount       *big.Int
				ImgIcon      string
			}{
				User:         v.User,
				Soliciter:    v.Soliciter,
				HospitalName: v.HospitalName,
				MedicalName:  v.MedicalName,
				Certificate:  fmt.Sprintf("0x%x", v.Certificate),
				Amount:       r2[k],
				ImgIcon:      g.ImgUrl,
			}
			r = append(r, r1)
		}

	}

	return r, err
}
