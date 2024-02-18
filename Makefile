# command to build proto files 
update-proto:
	protoc --go_out=./pkg --go_opt=paths=source_relative --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative api/proto/v1/address-api.proto

.PHONY: build run compose-infra-down compose-infra-up

config-local:
	./config.sh address-api local

config-dev:
	./config.sh address-api dev

config-prod:
	./config.sh address-api prod

run-api:
	go run ./cmd/address-api

run-api-local: config-local run-api

run-api-dev: config-dev run-api

run-api-prod: config-prod run-api