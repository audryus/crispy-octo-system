build:
	@go build -o bin/app

run: build 
	@./bin/app

test:
	@go test -v ./tests/...

docker:
	docker build -f ./build/docker/Dockerfile -t audryus/crispy-octo-system .
