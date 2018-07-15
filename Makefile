REG := lp150189
APP := mei
VERSION ?= $(shell cat VERSION)

all: build push 

build: 
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o mei main.go
	@docker build --build-arg VERSION=$(VERSION) -t $(REG)/$(APP):$(VERSION) --file $(APP).dockerfile .

push:
	@docker push $(REG)/$(APP):$(VERSION)

