package models

type Customer struct {
	ID          int     `json:"id"`
	NIK         string  `json:"nik"`
	FullName    string  `json:"full_name"`
	LegalName   string  `json:"legal_name"`
	Birthplace  string  `json:"birthplace"`
	Birthdate   string  `json:"birthdate"`
	Salary      float64 `json:"salary"`
	KTPPhoto    string  `json:"ktp_photo"`
	SelfiePhoto string  `json:"selfie_photo"`
	Limit       float64 `json:"limits"`
}
