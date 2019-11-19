all: clean

clean:
	find . -name 'inventory' -type f -exec rm -f {} \;
	go clean

run:
	go run src/inventory/server.go

main:
	go run src/inventory/main.go

format:
	go fmt goapi

install:
	go install inventory

build:
	go build inventory

start:
	./inventory

test-ping:
	curl localhost:3000/ping

docker-build:
	docker build -t jojojoseph09/go-inventory .
	docker images

docker-run:
	docker run --name go-inventory -td -p 3000:3000 jojojoseph09/go-inventory
	docker ps

docker-push:
	docker push jojojoseph09/go-inventory
