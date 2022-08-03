//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	adapter2 "github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/adapter"
	"github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/domain/adapter"
	"github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/domain/service"
)

func BuildAppContainer() (*adapter.TransactionHttpAdapter, error) {
	wire.Build(adapter2.NewLedgerService,
		wire.Bind(new(service.Ledger), new(*adapter2.LedgerService)),
		adapter.NewTransactionHttpAdapter,
	)
	return nil, nil
}
