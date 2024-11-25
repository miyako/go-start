# go-start
CLI to launch 4D CLI

work in progress; idea is to make an alternative to *Updater.exe*

## Usage

```
./start -applicationPath /Applications/4D\ 20.5/4D.app -project /Users/miyako/Desktop/example.4DLink
start.exe  -applicationPath "C:\Users\Public\Desktop\4D.lnk" -project "C:\Users\miyako\AppData\Roaming\4D\Favorites v20\Local\example.4DLink"
```

## Go Build

```
GOOS=darwin GOARCH=arm64 go build -o start-arm main.go
GOOS=darwin GOARCH=amd64 go build -o start-amd main.go
lipo -create start-arm start-amd -output start
GOOS=windows GOARCH=amd64 go build -o start.exe main.go
```
