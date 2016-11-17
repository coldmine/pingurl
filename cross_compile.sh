GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=`git describe --tags`" -o ./bin/urlcheck_linux_x86_64
GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=`git describe --tags`" -o ./bin/urlcheck_macos_x86_64
GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=`git describe --tags`" -o ./bin/urlcheck_windows_x86_64.exe
