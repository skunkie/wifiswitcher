commit := $(shell git rev-parse --short=8 HEAD)
LDFLAGS := "-s -w -X 'main.commit=$(commit)' -H windowsgui"

build:
	GOOS=windows GOARCH=amd64 go build -ldflags $(LDFLAGS) -trimpath -o ./bin/wifiswitcher.exe .

prepare:
	GO111MODULE=off go get -u github.com/josephspurrier/goversioninfo/cmd/goversioninfo
	goversioninfo

clean:
	@rm -rf ./bin

.PHONY: build clean prepare
