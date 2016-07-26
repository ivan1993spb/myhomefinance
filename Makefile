
install:

	bower install

build:

	go-bindata -o static_files.go -nometadata static/* static/libs/*

test:

	go test -v ./...

clean:

	git clean -f
	rm myhomefinance -f
	rm *.db *.db-journal -f
	rm static/libs -rf
