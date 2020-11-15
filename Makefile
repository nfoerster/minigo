full: linux windows
linux:
	go build -o .build/minigo main.go
windows:
	GOOS=windows GOARCH=amd64 go build -o .build/minigo.exe main.go
