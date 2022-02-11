export GO111MODULE := on

.PHONY: build install clean

cmd/vaws/vaws: *.go cmd/vaws/*.go go.*
	go build -o cmd/vaws/vaws -trimpath

test:
	cd cmd/vaws && go test -v

install: cmd/vaws/vaws
	install cmd/vaws/vaws /usr/local/bin/vaws

build:
	go build -o cmd/vaws/vaws main.go && chmod 755 cmd/vaws/vaws

clean:
	rm -f cmd/vaws/vaws
