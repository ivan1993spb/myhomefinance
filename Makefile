
test:

	go test -v ./...

clean:

	git clean -f
	rm myhomefinance -f
	rm *.db *.db-journal -f
