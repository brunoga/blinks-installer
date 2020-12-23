env GOARCH=amd64

env GOOS=linux go build -o blinks-installer-linux-amd64
env GOOS=darwin go build -o blinks-installer-darwin-amd64
env GOOS=windows go build -o blinks-installer-windows-amd64.exe

