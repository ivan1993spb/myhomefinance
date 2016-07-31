
deps:

install:

	npm install
	bower install

build:

	gulp build
	go generate -v
	go build -v

test:

	go test -v ./sqlite3mappers

clean:

	git clean -f
	rm -rf \
		myhomefinance \
		bindata_assetfs.go \
		*.db \
		*.db-journal \
		static/dist \
		bower_components \
		node_modules
