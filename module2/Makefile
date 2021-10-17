export tag=v1.0
root:
	export ROOT=github.com/zhouqi/CloudNativeHomework
build:
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64/httpserver .
release: build
	docker build -t zhouqi/httpserver:${tag} .
push: release
	docker push zhouqi/httpserver:${tag}
        
