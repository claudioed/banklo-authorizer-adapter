package adapter

import "github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/domain"

type LedgerService struct {
}

func (ls *LedgerService) Register(t *domain.Transaction) error {
	return nil
}

func (ls *LedgerService) Cancel(t *domain.Transaction) error {
	return nil
}

func (ls *LedgerService) Reversal(t *domain.Transaction) error {
	return nil
}

func NewLedgerService() *LedgerService {
	return &LedgerService{}
}
