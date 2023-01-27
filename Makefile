integration-tests:
	go build -o test/main cmd/main.go
	go test test/challenge_test.go -v

unit-tests:
	go test internal/handler/* -v
	go test internal/service/* -v