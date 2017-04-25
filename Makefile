
test:

	go list ./... | grep -v vendor | xargs go test -v
