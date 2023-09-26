package dto

type WithdrawalReq struct {
	NoRekening string `json:"no_rekening"`
	Nominal float64 `json:"nominal"`
}