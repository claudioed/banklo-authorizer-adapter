package client

import (
	api "github.com/modern-apis-architecture/banklo-authorizer-adapter/internal/ledger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func ProvideLedgerServiceClient() api.LedgerServiceClient {
	cc, err := grpc.Dial(os.Getenv("LEDGER_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	lsc := api.NewLedgerServiceClient(cc)
	return lsc
}
