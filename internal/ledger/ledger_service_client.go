package ledger

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

func ProvideLedgerServiceClient() LedgerServiceClient {
	cc, err := grpc.Dial(os.Getenv("LEDGER_URL"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	lsc := NewLedgerServiceClient(cc)
	return lsc
}
