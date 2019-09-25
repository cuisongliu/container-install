all: generator build test-docker test-containerd
generator-docker:
	export GO111MODULE="on" && go get &&  go test -v -run TestDocker_fetch github.com/cuisongliu/container-install
generator:
	go get -u github.com/a-urth/go-bindata/...
	go-bindata -pkg command -o install/command/assert.go install/command/
build:
	export GO111MODULE="on" && go get && go build -o container-install

test-docker:
	./container-install print -d

test-containerd:
	./container-install print
