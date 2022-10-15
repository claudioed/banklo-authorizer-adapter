package adapter

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/api"
	"github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/domain"
	"github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/domain/service"
	"time"
)

type TransactionHttpAdapter struct {
	ledger service.Ledger
}

func NewTransactionHttpAdapter(ledger service.Ledger) *TransactionHttpAdapter {
	return &TransactionHttpAdapter{ledger: ledger}
}

func (tha *TransactionHttpAdapter) CreateTransaction(ctx echo.Context) error {
	t := &api.RequestTransaction{}
	if err := ctx.Bind(t); err != nil {
		return err
	}
	tr := &domain.Transaction{
		Id:                *t.TransactionData.TransactionId,
		Amount:            float64(*t.TransactionData.Amount),
		LocalTime:         time.Time{},
		CardId:            *t.TransactionData.CardId,
		MerchantId:        *t.MerchantCode,
		CurrencyCode:      *t.CurrencyCode,
		CountryCode:       *t.CountryCode,
		AuthorizationCode: *t.AuthorizationCode,
	}
	nc := context.WithValue(ctx.Request().Context(), "external-auth", ctx.Request().Header["Authorization"][0])
	err := tha.ledger.Register(nc, tr)
	if err != nil {
		return err
	}
	return ctx.JSON(201, api.ResponseTransactions{
		RegisteredAt:  time.Now(),
		TransactionId: tr.Id,
	})
}

func (tha *TransactionHttpAdapter) RequestCancellation(ctx echo.Context, id api.Id) error {
	return nil
}

func (tha *TransactionHttpAdapter) RequestReversal(ctx echo.Context, id api.Id) error {
	return nil
}
