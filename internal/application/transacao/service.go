package application

import (
	"context"

	"github.com/joaomarcosbc/rinha-backend-2024/internal/domain"
	infra "github.com/joaomarcosbc/rinha-backend-2024/internal/infra/transacao"
)

type service struct {
	repository infra.Repository
}

func NewService(r infra.Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CriarTransacao(ctx context.Context, trasacao *domain.Transacao) (domain.TransacaoResposta, error) {
	panic("unimplemented")
}

// Ideally "business logic" would be implemented in this method
func (s *service) ObterExtrato(ctx context.Context, id int) (domain.ExtratoDto, error) {
	extrato, err := s.repository.ObterExtrato(ctx, id)

	return extrato, err
}
