all: linux windows
linux:
	go build -o .build/minigo cmd/minigo/main.go
windows:
	GOOS=windows GOARCH=amd64 go build -o .build/minigo.exe cmd/minigo/main.go
