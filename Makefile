all: generator build test
generator:
	export GO111MODULE="disable"  && go get -u github.com/a-urth/go-bindata/...
	go-bindata -pkg command -o install/command/assert.go install/command/
build:
	export GO111MODULE="on" && go get && go build -o container-install

test:
	./container-install print -d
