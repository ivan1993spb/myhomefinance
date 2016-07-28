
deps:

install:

	npm install
	bower install

build:

	gulp build

	go-bindata-assetfs -nometadata \
		-ignore "\.bower\.json" \
		-ignore "static/src/" \
		static/...

	go build -v

test:

	go test -v ./sqlite3mappers

clean:

	git clean -f
	rm myhomefinance -f
	rm *.db *.db-journal -f
	rm static/libs static/dist -rf
	rm node_modules -rf
