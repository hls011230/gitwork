package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"A11Smile/eth"
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func CreateCertificate(uid int, meta model.PostCertificate) error {

	DB := db.Get()
	var user model.Wallet
	DB.Table("users").First(&user, "id = ?", uid)

	nonce, err := eth.Client.PendingNonceAt(context.Background(), common.HexToAddress(user.BlockAddress))
	if err != nil {

		return err
	}

	privateKey, err := crypto.HexToECDSA(user.PrivateKey)
	if err != nil {

		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, eth.ChainID)
	if err != nil {
		return err
	}

	auth.GasPrice = eth.GasPrice
	auth.GasLimit = uint64(6000000)
	auth.Nonce = big.NewInt(int64(nonce))

	_, err = eth.Ins.SetCertificate(auth, meta.ArrayMedicalHistory, meta.ArrayMedicalExaminationReport, meta.CertificateName)
	if err != nil {
		return err
	}

	return nil
}

func ShowAllCertificate(uid int) (interface{}, error) {
	DB := db.Get()
	var user model.Wallet
	DB.Table("users").First(&user, "id = ?", uid)

	res, err := eth.Ins.AllUserSerialNames(&bind.CallOpts{Context: context.Background(), From: common.HexToAddress(user.BlockAddress)})
	if err != nil {
		return nil, err
	}

	return res, nil

}

func ShowDetailsCertificate(uid int, n string) (interface{}, error) {
	DB := db.Get()
	var user model.Wallet
	DB.Table("users").First(&user, "id = ?", uid)
	hash, err := eth.Ins.UserSerialName(nil, common.HexToAddress(user.BlockAddress), n)

	r, err := eth.Ins.ViewUserCertificate(nil, hash)
	if err != nil {
		return "", err
	}

	t := r.Time.String()
	t2, _ := strconv.Atoi(t)
	tm := time.Unix(int64(t2), 0)

	res := struct {
		Address                     string `json:"address"`
		MedicalHistoryNum           int    `json:"medical_history_num"`
		MedicalExaminationReportNum int    `json:"medical_examination_report_num"`
		BlockNum                    string `json:"block_num"`
		Serial                      string `json:"serial"`
		BlockTime                   string `json:"block_time"`
	}{
		Address:                     r.User.String(),
		MedicalHistoryNum:           len(r.M1),
		MedicalExaminationReportNum: len(r.M2),
		BlockNum:                    r.BlockNum.String(),
		Serial:                      fmt.Sprintf("0x%x", r.Serial),
		BlockTime:                   fmt.Sprintf(tm.Format("2006-01-02 15:04:05")),
	}

	return res, nil

}
