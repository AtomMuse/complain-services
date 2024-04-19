run-complain:
	go run cmd/complain/main.go

test-coverage:
	mkdir -p coverage
	go test -race -short -v -coverprofile coverage/cover.out ./...
	go tool cover -html=coverage/cover.out

gen-swag:
	swag init -d ./cmd/complain,./handler/complainhandler -o ./cmd/complain/doc --pd