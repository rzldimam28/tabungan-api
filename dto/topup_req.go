package dto

type TopUpReq struct {
	NoRekening string `json:"no_rekening"`
	Nominal float64 `json:"nominal"`
}