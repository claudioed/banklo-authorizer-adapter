package service

import (
	"context"
	"github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/domain"
)

type Ledger interface {
	Register(ctx context.Context, t *domain.Transaction) error
	Cancel(ctx context.Context, t *domain.Transaction) error
	Reversal(ctx context.Context, t *domain.Transaction) error
}
