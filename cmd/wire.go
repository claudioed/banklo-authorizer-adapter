//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/adapter"
	"github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/domain/service"
	"github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/ledger/client"
)

func BuildAppContainer() (*adapter.TransactionHttpAdapter, error) {
	wire.Build(client.ProvideLedgerServiceClient, client.NewInternalLedgerService,
		wire.Bind(new(service.Ledger), new(*client.InternalLedgerService)),
		adapter.NewTransactionHttpAdapter,
	)
	return nil, nil
}
