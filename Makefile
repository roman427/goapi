build-goapi:
	GOOS=linux GOARCH=amd64 go build -mod=vendor -v -o goapi ./cmd/goapi