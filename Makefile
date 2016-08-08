
install: deps test build

	go install

deps:

	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/elazarl/go-bindata-assetfs/...
	npm install

test:

	go test -v ./sqlite3mappers

build:

	gulp build
	go generate -v
	go build -v

clean:

	git clean -f
	rm -rf \
		myhomefinance \
		bindata_assetfs.go \
		*.db \
		*.db-journal \
		static/dist \
		node_modules

start:

	./myhomefinance &
	gulp watch &

stop:

	kill -9 `pidof myhomefinance`
	kill -9 `ps -A -o pid,cmd | grep "gulp" | head -n 1 | awk '{print $$1}'`
