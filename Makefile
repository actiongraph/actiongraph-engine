LOCAL_NAME ?= actiongraph_engine:latest

clean:
	@echo "cleaning .."
	rm -rf build/
	@echo "cleaned"

build: clean
	@echo "begin building .."
	go build -o build/app main.go
	mkdir ./build/data
	@echo "building done"

build-linux: clean
	@echo "begin building .."
	GOOS=linux GOARCH=amd64 go build -o build/app main.go
	@echo "building done"

run: build
	@clear
	@./build/app

ensure:
	dep ensure

gen-protobuf:
	@echo "compiling protobuf definitions .."
	protoc --go_out=plugins=grpc:./ proto/scrapper.proto
	@echo "protobuf definitions compiled"

docker: build-linux
	@echo "building docker image .."
	docker rmi ${LOCAL_NAME} -f
	docker build . -t ${LOCAL_NAME}
	@echo "building done"