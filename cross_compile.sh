GOOS=linux GOARCH=amd64 go build -o ./bin/linux/pingurl pingurl.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/osx/pingurl pingurl.go
GOOS=windows GOARCH=amd64 go build -o ./bin/windows/pingurl.exe pingurl.go
