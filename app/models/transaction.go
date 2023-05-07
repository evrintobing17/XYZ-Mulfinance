package models

type Transaction struct {
	ID          int     `json:"id"`
	ContractNo  string  `json:"contract_number"`
	OTR         int     `json:"otr"`
	AdminFee    int     `json:"admin_fee"`
	Installment int     `json:"installment"`
	Interest    float64 `json:"interest"`
	AssetName   string  `json:"asset_name"`
	CustomerID  int     `json:"customer_id"`
	Limits      []Limit `json:"limits"`
}
