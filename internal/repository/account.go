package repository

import (
	"context"
	"database/sql"

	"github.com/rzldimam28/tabungan-api/domain"
)

var (
	selectAccountByNoRekQuery = "SELECT a.id, a.nama, a.nik, a.no_hp, a.no_rekening, a.saldo FROM accounts a WHERE a.no_rekening = $1"
	insertAccountQuery = "INSERT INTO accounts (nama, nik, no_hp, no_rekening, saldo) VALUES ($1, $2, $3, $4, $5)"
	updateAccountQuery = "UPDATE accounts SET saldo = $1 WHERE id = $2"
)

type accountRepository struct {
}

func NewAccount() domain.AccountRepository {
	return &accountRepository{}	
}

func (ths *accountRepository) FindByNoRekening(ctx context.Context, tx *sql.Tx, noRekening string) (*domain.Account, error) {
	rows, err := tx.QueryContext(ctx, selectAccountByNoRekQuery, noRekening)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var account domain.Account
	if rows.Next() {
		err = rows.Scan(&account.ID, &account.Nama, &account.Nik, &account.NoHp, &account.NoRekening, &account.Saldo)
		if err != nil {
			return nil, err
		}
		return &account, nil
	} else {
		return nil, domain.ErrAccountNotFound
	}
}

func (ths *accountRepository) Insert(ctx context.Context, tx *sql.Tx, account domain.Account) error {
	_, err := tx.ExecContext(ctx, insertAccountQuery, account.Nama, account.Nik, account.NoHp, account.NoRekening, account.Saldo)
	if err != nil {
		return err
	}

	return nil
}

func (ths *accountRepository) Update(ctx context.Context, tx *sql.Tx, account domain.Account) error {
	_, err := tx.ExecContext(ctx, updateAccountQuery, account.Saldo, account.ID)
	if err != nil {
		return err
	}
	return nil
}