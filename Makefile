stubs:
	 oapi-codegen --package api api/openapi.yaml  > internal/api/authorizer_adapter.gen.go
ledger-client:
	protoc -I ./api --go_out ./internal/ledger --go_opt paths=source_relative --go-grpc_out ./internal/ledger --go-grpc_opt paths=source_relative ./api/ledger.proto