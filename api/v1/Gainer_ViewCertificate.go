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
)

func Gainer_ViewCertificate(gid int, medicalName string) ([]interface{}, error) {
	DB := db.Get()
	var w model.Wallet
	DB.Table("gainers").First(&w, "id = ?", gid)
	res, err := eth.Ins.GainerSeeuserUploadMedical(&bind.CallOpts{Context: context.Background(), From: common.HexToAddress(w.BlockAddress)})
	var r []interface{}
	for _, v := range res {
		if v.MedicalName == medicalName {

			r1 := struct {
				User         common.Address
				State        bool
				Soliciter    common.Address
				HospitalName string
				MedicalName  string
				Certificate  string
				Amount       *big.Int
			}{
				User:         v.User,
				Soliciter:    v.Soliciter,
				HospitalName: v.HospitalName,
				MedicalName:  v.MedicalName,
				Certificate:  fmt.Sprintf("0x%x", v.Certificate),
			}
			r = append(r, r1)
		}
	}
	return r, err

}

func ShowUserCertificateDetails(serial string) (interface{}, error) {
	r, err := eth.Ins.ViewUserCertificate(nil, common.HexToHash(serial))
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
