package infra

import (
	"context"
	"time"

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
	cliente, err := r.ObterLimiteESaldo(ctx, id)
	if err != nil {
		return domain.ExtratoDto{}, err
	}

	query := `SELECT valor, tipo, descricao, realizada_em 
			  FROM transacoes 
			  WHERE cliente_id = $1 
			  ORDER BY realizada_em DESC LIMIT 10`

	rows, err := r.db.Query(ctx, query, id)

	if err != nil {
		return domain.ExtratoDto{}, err
	}

	var ultimasTransacoes []domain.TransacaoDto

	for rows.Next() {
		t := domain.TransacaoDto{}
		if err = rows.Scan(&t.Valor, &t.Tipo, &t.Descricao, &t.RealizadaEm); err != nil {
			return domain.ExtratoDto{}, err
		}
		ultimasTransacoes = append(ultimasTransacoes, t)
	}

	extrato := domain.ExtratoDto{
		Saldo: domain.SaldoClienteDto{
			Total:       cliente.Saldo,
			DataExtrato: time.Now(),
			Limite:      cliente.Limite,
		},
		UltimasTransacoes: ultimasTransacoes,
	}

	return extrato, err
}

func (r *repository) ObterLimiteESaldo(ctx context.Context, id int) (domain.Cliente, error) {
	query := "SELECT limite, saldo from clientes WHERE id=$1;"
	row := r.db.QueryRow(ctx, query, id)
	cliente := domain.Cliente{}

	if err := row.Scan(&cliente.Limite, &cliente.Saldo); err != nil {
		return domain.Cliente{}, err
	}

	return cliente, nil
}
