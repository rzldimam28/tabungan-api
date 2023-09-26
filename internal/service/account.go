package service

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/rzldimam28/tabungan-api/domain"
	"github.com/rzldimam28/tabungan-api/dto"
	"github.com/rzldimam28/tabungan-api/internal/util"
)

type accountService struct {
	db *sql.DB
	accountRepository domain.AccountRepository
	mutationRepository domain.MutationRepository
}

func NewAccount(db *sql.DB, accountRepository domain.AccountRepository, mutationRepository domain.MutationRepository) domain.AccountService {
	return &accountService{
		db: db,
		accountRepository: accountRepository,
		mutationRepository: mutationRepository,
	}
}

func (ths *accountService) Create(ctx context.Context, req dto.RegisterAccountReq) (*dto.RegisterAccountRes, error) {
	tx, err := ths.db.Begin()
	if err != nil {
		return nil, err
	}

	noRek := util.GenerateRandomNumber(10)
	account := domain.Account{
		Nama: req.Nama,
		Nik: req.Nik,
		NoHp: req.NoHp,
		NoRekening: noRek,
		Saldo: float64(0),
	}
	err = ths.accountRepository.Insert(ctx, tx, account)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &dto.RegisterAccountRes{
		NoRekening: noRek,
	}, nil
}

func (ths *accountService) TopUp(ctx context.Context, req dto.TopUpReq) (*dto.TopUpRes, error) {
	tx, err := ths.db.Begin()
	if err != nil {
		return nil, err
	}

	account, err := ths.accountRepository.FindByNoRekening(ctx, tx, req.NoRekening)
	if err != nil {
		log.Println("here", err)
		tx.Rollback()
		if err == sql.ErrNoRows {
			return nil, domain.ErrAccountNotFound
		}
		return nil, err
	}

	account.Saldo += req.Nominal

	err = ths.accountRepository.Update(ctx, tx, *account)
	if err != nil {
		log.Println("here 2", err)
		tx.Rollback()
		return nil, err
	}

	mutation := domain.Mutation{
		NoRekening: account.NoRekening,
		KodeTransaksi: "C",
		Nominal: req.Nominal,
		CreatedAt: time.Now(),
	}
	err = ths.mutationRepository.Insert(ctx, tx, mutation)
	if err != nil {
		log.Println("here 3", err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &dto.TopUpRes{
		Saldo: account.Saldo,
	}, nil
}

func (ths *accountService) Withdrawal(ctx context.Context, req dto.WithdrawalReq) (*dto.WithdrawalRes, error) {
	tx, err := ths.db.Begin()
	if err != nil {
		return nil, err
	}

	account, err := ths.accountRepository.FindByNoRekening(ctx, tx, req.NoRekening)
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			return nil, domain.ErrAccountNotFound
		}
		return nil, err
	}

	if account.Saldo < req.Nominal {
		tx.Rollback()
		return nil, domain.ErrNotEnoughBalance
	}

	account.Saldo -= req.Nominal

	err = ths.accountRepository.Update(ctx, tx, *account)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	mutation := domain.Mutation{
		NoRekening: account.NoRekening,
		KodeTransaksi: "D",
		Nominal: req.Nominal,
		CreatedAt: time.Now(),
	}
	err = ths.mutationRepository.Insert(ctx, tx, mutation)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &dto.WithdrawalRes{
		Saldo: account.Saldo,
	}, nil
}

func (ths *accountService) GetByNoRekening(ctx context.Context, noRekening string) (*dto.CheckBalanceAccountRes, error) {
	tx, err := ths.db.Begin()
	if err != nil {
		return nil, err
	}

	account, err := ths.accountRepository.FindByNoRekening(ctx, tx, noRekening)
	if err != nil {
		tx.Rollback()
		if err == sql.ErrNoRows {
			return nil, domain.ErrAccountNotFound
		}
		return nil, err
	}

	return &dto.CheckBalanceAccountRes{
		Saldo: account.Saldo,
	}, nil
}