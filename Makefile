BIN := dbloader
MAIN := cmd/dbloader/main.go

mac: $(MAIN)
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BIN) $(MAIN) 

linux: $(MAIN)
	GOOS=linux GOARCH=amd64 go build -o bin/$(BIN) $(MAIN)

windows: $(MAIN)
	GOOS=windows GOARCH=amd64 go build -o bin/$(BIN) $(MAIN)
