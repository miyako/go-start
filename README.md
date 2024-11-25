# go-start
CLI to launch 4D CLI

## Go Build

```
GOOS=darwin GOARCH=arm64 go build -o start-arm main.go
GOOS=darwin GOARCH=amd64 go build -o start-amd main.go
lipo -create start-arm start-amd -output start
GOOS=windows GOARCH=amd64 go build -o start.exe main.go
```
