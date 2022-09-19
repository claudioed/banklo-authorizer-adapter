package ledger

import (
	"context"
	"github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type InternalLedgerService struct {
	ledgerCli LedgerServiceClient
}

func (ls *InternalLedgerService) Register(ctx context.Context, t *domain.Transaction) error {
	rc := &RequestConfirmation{
		Transaction: &ConfirmationTransactionData{
			ExternalId:        t.Id,
			Value:             float32(t.Amount),
			LocalTime:         timestamppb.New(t.LocalTime),
			MerchantId:        t.MerchantId,
			CardId:            t.CardId,
			CountryCode:       t.CountryCode,
			CurrencyCode:      t.CurrencyCode,
			AuthorizationCode: t.AuthorizationCode,
		},
		Type: TransactionType_CONFIRMATION,
	}
	_, err := ls.ledgerCli.Confirmation(ctx, rc)
	if err != nil {
		return err
	}
	return nil
}

func (ls *InternalLedgerService) Cancel(ctx context.Context, t *domain.Transaction) error {
	rc := &RequestConfirmation{
		Transaction: &ConfirmationTransactionData{
			ExternalId: t.Id,
			Value:      float32(t.Amount),
			LocalTime:  timestamppb.New(t.LocalTime),
			MerchantId: t.MerchantId,
			CardId:     t.CardId,
		},
		Type: TransactionType_CANCELLATION,
	}
	_, err := ls.ledgerCli.Confirmation(ctx, rc)
	if err != nil {
		return err
	}
	return nil
}

func (ls *InternalLedgerService) Reversal(ctx context.Context, t *domain.Transaction) error {
	rc := &RequestConfirmation{
		Transaction: &ConfirmationTransactionData{
			ExternalId: t.Id,
			Value:      float32(t.Amount),
			LocalTime:  timestamppb.New(t.LocalTime),
			MerchantId: t.MerchantId,
			CardId:     t.CardId,
		},
		Type: TransactionType_REVERSAL,
	}
	_, err := ls.ledgerCli.Confirmation(ctx, rc)
	if err != nil {
		return err
	}
	return nil
}

func NewInternalLedgerService(ledgerCli LedgerServiceClient) *InternalLedgerService {
	return &InternalLedgerService{ledgerCli: ledgerCli}
}
