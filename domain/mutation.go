package domain

import (
	"context"
	"database/sql"
	"time"

	"github.com/rzldimam28/tabungan-api/dto"
)

type Mutation struct {
	ID            int64     `db:"id"`
	NoRekening    string    `db:"no_rekening"`
	KodeTransaksi string    `db:"kode_transaksi"`
	Nominal       float64   `db:"nominal"`
	CreatedAt     time.Time `db:"created_at"`
}

type MutationRepository interface {
	FindByNoRekening(ctx context.Context, tx *sql.Tx, noRekening string) ([]*Mutation, error)
	Insert(ctx context.Context, tx *sql.Tx, mutation Mutation) error
}

type MutationService interface {
	GetByNoRekening(ctx context.Context, noRekening string) ([]*dto.MutationRes, error)
}
