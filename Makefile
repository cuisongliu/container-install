
generator:
	go get -u github.com/a-urth/go-bindata/...
	go-bindata -pkg command -o install/command/assert.go install/command/
