package service

import "github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/domain"

type Ledger interface {
	Register(t *domain.Transaction) error

	Cancel(t *domain.Transaction) error

	Reversal(t *domain.Transaction) error
}
