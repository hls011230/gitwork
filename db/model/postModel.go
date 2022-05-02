package model

type PostEmail struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type PostUserFile struct {
	FileName string `json:"file_name"`
}

type PostDetails struct {
	Address     string `json:"address"`
	MedicalName string `json:"medical_name"`
}

type PostCertificate struct {
	ArrayMedicalHistory           []string `json:"array_medical_history"`
	ArrayMedicalExaminationReport []string `json:"array_medical_examination_report"`
}

type PostCertificateHash struct {
	Serial string `json:"serial"`
}
