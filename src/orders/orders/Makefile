all: clean

clean:
	find . -name 'orders' -type f -exec rm -f {} \;
	go clean

run:
	go run src/app/server.go

main:
	go run src/app/main.go

format:
	go fmt orders

install:
	go install orders

build:
	go build orders

start:
	./orders

test-ping:
	curl localhost:3000/ping

docker-build:
	docker build -t harshrajm/orders .
	docker images

docker-run:
	docker run --name orders -td -p 3000:3000 harshrajm/orders
	docker ps

docker-push:
	docker push harshrajm/orders
