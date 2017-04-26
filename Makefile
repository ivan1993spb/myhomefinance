
test:

	go list ./... | grep -v vendor | xargs go vet -v
	go list ./... | grep -v vendor | xargs go test -v
