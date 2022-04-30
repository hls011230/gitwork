package v1

import (
	"A11Smile/db/model"
	"A11Smile/eth"
	"A11Smile/eth/gen"

	"github.com/ethereum/go-ethereum/common"
)

func DisplayHomepage() ([]gen.UploadMedicalrecords_gainergainer_upMedicalInformation, error) {
	res, err := eth.Ins.SeeGainerMedicalInformationsName(nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ShowDetailsPage(details model.PostDetails) (gen.UploadMedicalrecords_gainergainer_upMedicalInformation1, error) {
	res, err := eth.Ins.SeeGainerMedicalInformations(nil, common.HexToAddress(details.Address), details.MedicalName)
	if err != nil {
		return res, err
	}
	return res, nil
}
