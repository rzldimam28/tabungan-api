package repository

import (
	"context"
	"database/sql"

	"github.com/rzldimam28/tabungan-api/domain"
)

var (
	selectMutationByNoRekQuery = "SELECT m.id, m.no_rekening, m.kode_transaksi, m.nominal, m.created_at FROM mutations m WHERE m.no_rekening = $1"
	insertMutationQuery = "INSERT INTO mutations (no_rekening, kode_transaksi, nominal, created_at) VALUES ($1, $2, $3, $4)"
)

type mutationRepository struct {
}

func Newmutation() domain.MutationRepository {
	return &mutationRepository{}
}

func (ths *mutationRepository) FindByNoRekening(ctx context.Context, tx *sql.Tx, noRekening string) ([]*domain.Mutation, error) {
	rows, err := tx.QueryContext(ctx, selectMutationByNoRekQuery, noRekening)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var mutations []*domain.Mutation
	for rows.Next() {
		mutation := &domain.Mutation{}
		err = rows.Scan(&mutation.ID, &mutation.NoRekening, &mutation.KodeTransaksi, &mutation.Nominal, &mutation.CreatedAt)
		if err != nil {
			return nil, err
		}
		mutations = append(mutations, mutation)
	}

	return mutations, nil
}

func (ths *mutationRepository) Insert(ctx context.Context, tx *sql.Tx, mutation domain.Mutation) error {
	_, err := tx.ExecContext(ctx, insertMutationQuery, mutation.NoRekening, mutation.KodeTransaksi, mutation.Nominal, mutation.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}