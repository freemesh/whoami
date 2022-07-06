VERSION=$(shell git describe --tags --always)

.PHONY: docker.build
docker.build:
	docker build --build-arg APP_VERSION=$(VERSION) -t freemesh/whoami:$(VERSION) .

.PHONY: docker.push
docker.push:
	docker push freemesh/whoami:$(VERSION)

.PHONY: docker.all
docker.all:docker.build docker.push

.PHONY: build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/whoami ./main.go

.PHONY: run
run:
	go run -ldflags "-X main.Version=$(VERSION)" ./main.go
