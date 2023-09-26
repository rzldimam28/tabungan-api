package service

import (
	"context"
	"database/sql"

	"github.com/rzldimam28/tabungan-api/domain"
	"github.com/rzldimam28/tabungan-api/dto"
)

type mutationService struct {
	db                 *sql.DB
	mutationRepository domain.MutationRepository
}

func NewMutation(db *sql.DB, mutationRepository domain.MutationRepository) domain.MutationService {
	return &mutationService{
		db:                db,
		mutationRepository: mutationRepository,
	}
}

func (ths *mutationService) GetByNoRekening(ctx context.Context, noRekening string) ([]*dto.MutationRes, error) {
	tx, err := ths.db.Begin()
	if err != nil {
		return nil, err
	}

	mutations, err := ths.mutationRepository.FindByNoRekening(ctx, tx, noRekening)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if len(mutations) <= 0 {
		tx.Rollback()
		return nil, domain.ErrMutationsNotFound
	}

	var mutationResps []*dto.MutationRes
	for _, mutation := range mutations {
		mutationResp := &dto.MutationRes{
			KodeTransaksi: mutation.KodeTransaksi,
			Nominal: mutation.Nominal,
			Waktu: mutation.CreatedAt,
		}
		mutationResps = append(mutationResps, mutationResp)
	}

	return mutationResps, nil
}