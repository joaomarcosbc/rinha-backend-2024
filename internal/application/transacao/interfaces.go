package application

import (
	"context"

	"github.com/joaomarcosbc/rinha-backend-2024/internal/domain"
)

type Service interface {
	CriarTransacao(ctx context.Context, trasacao *domain.Transacao) (domain.TransacaoResposta, error)
	ObterExtrato(ctx context.Context, id int) (domain.ExtratoDto, error)
}
