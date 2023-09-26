package domain

import (
	"context"
	"database/sql"

	"github.com/rzldimam28/tabungan-api/dto"
)

type Account struct {
	ID   int64  `db:"id"`
	Nama string `db:"nama"`
	Nik  string `db:"nik"`
	NoHp string `db:"no_hp"`
	NoRekening string `db:"no_rekening"`
	Saldo float64 `db:"saldo"`
}

type AccountRepository interface {
	FindByNoRekening(ctx context.Context, tx *sql.Tx, noRekening string) (*Account, error)
	Insert(ctx context.Context, tx *sql.Tx, account Account) error
	Update(ctx context.Context, tx *sql.Tx, account Account) error
}

type AccountService interface {
	Create(ctx context.Context, req dto.RegisterAccountReq) (*dto.RegisterAccountRes, error)
	TopUp(ctx context.Context, req dto.TopUpReq) (*dto.TopUpRes, error)
	Withdrawal(ctx context.Context, req dto.WithdrawalReq) (*dto.WithdrawalRes, error)
	GetByNoRekening(ctx context.Context, noRekening string) (*dto.CheckBalanceAccountRes, error)
}