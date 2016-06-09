
generate:

	swagger generate server -f swagger.yaml
	go generate ./sqlite3mappers

