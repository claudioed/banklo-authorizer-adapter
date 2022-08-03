package adapter

import (
	"github.com/labstack/echo/v4"
	"github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/api"
	"github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/domain/service"
)

type TransactionHttpAdapter struct {
	ledger service.Ledger
}

func NewTransactionHttpAdapter(ledger service.Ledger) *TransactionHttpAdapter {
	return &TransactionHttpAdapter{ledger: ledger}
}

func (tha *TransactionHttpAdapter) CreateTransaction(ctx echo.Context) error {
	return nil
}

func (tha *TransactionHttpAdapter) RequestCancellation(ctx echo.Context, id api.Id) error {
	return nil
}

func (tha *TransactionHttpAdapter) RequestReversal(ctx echo.Context, id api.Id) error {
	return nil
}
