GOOS=linux GOARCH=amd64 go build -o ./bin/linux_x86_64/urlcheck urlcheck.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/macos_x86_64/urlcheck urlcheck.go
GOOS=windows GOARCH=amd64 go build -o ./bin/windows_x86_64/urlcheck.exe urlcheck.go
