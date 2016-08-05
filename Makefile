
deps:

install:

	npm install

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
		node_modules

start:

	./myhomefinance &
	gulp watch &

stop:

	kill -9 `pidof myhomefinance`
	kill -9 `ps -A -o pid,cmd | grep "gulp" | head -n 1 | awk '{print $$1}'`
