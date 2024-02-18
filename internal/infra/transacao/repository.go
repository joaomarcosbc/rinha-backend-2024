package infra

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joaomarcosbc/rinha-backend-2024/internal/domain"
)

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CriarTransacao(ctx context.Context, transacao *domain.Transacao) (domain.TransacaoResposta, error) {
	panic("unimplemented")
}

func (r *repository) ObterExtrato(ctx context.Context, id int) (domain.ExtratoDto, error) {
	panic("unimplemented")
}

func (r *repository) ObterLimiteESaldo(ctx context.Context, id int) (domain.Cliente, error) {
	query := "SELECT limite, saldo from clientes WHERE id=$1;"
	row := r.db.QueryRow(ctx, query, id)
	cliente := domain.Cliente{}
	var limite, saldo int64
	err := row.Scan(&limite, &saldo)

	if err != nil {
		return cliente, err
	}

	cliente.Limite = limite
	cliente.Saldo = saldo

	return cliente, err
}
