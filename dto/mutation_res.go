package dto

import "time"

type MutationRes struct {
	KodeTransaksi string    `json:"kode_transaksi"`
	Nominal       float64   `json:"nominal"`
	Waktu         time.Time `json:"waktu"`
}
