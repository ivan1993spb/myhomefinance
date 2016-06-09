
generate:
	# generate api

	swagger generate server -f swagger.yaml

	# generate go code

	go generate ./sqlite3mappers

