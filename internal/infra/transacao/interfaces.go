package infra

import (
	"context"

	"github.com/joaomarcosbc/rinha-backend-2024/internal/domain"
)

type Repository interface {
	CriarTransacao(ctx context.Context, transacao *domain.Transacao) (domain.TransacaoResposta, error)
	ObterSaldo(ctx context.Context, id int) (domain.Cliente, error)
	ObterExtrato(ctx context.Context, id int) (domain.ExtratoDto, error)
}
