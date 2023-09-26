package dto

type RegisterAccountRes struct {
	NoRekening string `json:"no_rekening"`
}

type CheckBalanceAccountRes struct {
	Saldo float64 `json:"saldo"`
}