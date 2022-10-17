package client

import (
	"context"
	"github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/domain"
	api "github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/ledger"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type InternalLedgerService struct {
	ledgerCli api.LedgerServiceClient
}

func (ls *InternalLedgerService) Register(ctx context.Context, t *domain.Transaction) error {
	rc := &api.RequestConfirmation{
		Transaction: &api.ConfirmationTransactionData{
			ExternalId:        t.Id,
			Value:             float32(t.Amount),
			LocalTime:         timestamppb.New(t.LocalTime),
			MerchantId:        t.MerchantId,
			CardId:            t.CardId,
			CountryCode:       t.CountryCode,
			CurrencyCode:      t.CurrencyCode,
			AuthorizationCode: t.AuthorizationCode,
		},
		Type: api.TransactionType_CONFIRMATION,
	}
	_, err := ls.ledgerCli.Confirmation(ctx, rc)
	if err != nil {
		log.Errorf("error to send data to ledger. err %v", err)
		return err
	}
	return nil
}

func (ls *InternalLedgerService) Cancel(ctx context.Context, t *domain.Transaction) error {
	rc := &api.RequestConfirmation{
		Transaction: &api.ConfirmationTransactionData{
			ExternalId: t.Id,
			Value:      float32(t.Amount),
			LocalTime:  timestamppb.New(t.LocalTime),
			MerchantId: t.MerchantId,
			CardId:     t.CardId,
		},
		Type: api.TransactionType_CANCELLATION,
	}
	_, err := ls.ledgerCli.Confirmation(ctx, rc)
	if err != nil {
		return err
	}
	return nil
}

func (ls *InternalLedgerService) Reversal(ctx context.Context, t *domain.Transaction) error {
	rc := &api.RequestConfirmation{
		Transaction: &api.ConfirmationTransactionData{
			ExternalId: t.Id,
			Value:      float32(t.Amount),
			LocalTime:  timestamppb.New(t.LocalTime),
			MerchantId: t.MerchantId,
			CardId:     t.CardId,
		},
		Type: api.TransactionType_REVERSAL,
	}
	_, err := ls.ledgerCli.Confirmation(ctx, rc)
	if err != nil {
		return err
	}
	return nil
}

func NewInternalLedgerService(ledgerCli api.LedgerServiceClient) *InternalLedgerService {
	return &InternalLedgerService{ledgerCli: ledgerCli}
}
