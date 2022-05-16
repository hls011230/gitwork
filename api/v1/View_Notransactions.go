package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func ShowNoTransactionsHandler(id int) ([]interface{}, error) {
	// 根据uid获取用户的私钥和地址
	DB := db.Get()
	var w model.Wallet
	DB.Table("users").First(&w, "id = ?", id)
	r1, err := eth.Ins.UserYSeeCertificateState(&bind.CallOpts{Context: context.Background(), From: common.HexToAddress(w.BlockAddress)})

	var r []interface{}
	for _, v := range r1 {
		if v.MedicalName != "" {
			var g model.Gainer
			DB.Table("gainers").First(&g, "block_address = ?", v.Soliciter.String())
			r1 := struct {
				User         common.Address
				Soliciter    common.Address
				HospitalName string
				MedicalName  string
				Certificate  string
				ImgIcon      string
			}{
				User:         v.User,
				Soliciter:    v.Soliciter,
				HospitalName: v.HospitalName,
				MedicalName:  v.MedicalName,
				Certificate:  fmt.Sprintf("0x%x", v.Certificate),
				ImgIcon:      g.ImgUrl,
			}

			r = append(r, r1)
		}
	}

	return r, err
}
