GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

fmt:
	@echo "Running source files through gofmt..."
	gofmt -w $(GOFMT_FILES)

build:
	protoc -I ./cmd/grpc/routeguide/ ./cmd/grpc/routeguide/service.proto --go_out=plugins=grpc:./cmd/grpc/routeguide
docker-grpc:
	sudo docker build -t grpc-api -f ./cmd/grpc/Dockerfile .
docker-subscribers:
	sudo docker build -t subscribers-api -f ./cmd/subscribers/Dockerfile .
docker-api:
	sudo docker build -t http-api -f ./cmd/api/Dockerfile .
docker-up:
	sudo docker-compose build
	sudo docker-compose up
docker-upD:
	sudo docker-compose up -d
lint-all:
	golangci-lint run --enable-all
