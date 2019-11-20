
all: clean

clean: 
	find . -name 'payments' -type f -exec rm -f {} \; 
	go clean

go-get:
	rm -rf src/github.com
	go get -v github.com/codegangsta/negroni
	go get -v github.com/gorilla/mux
	go get -v github.com/unrolled/render
	go get -v github.com/satori/go.uuid

format:
	go fmt payments

install:
	go install payments

build:
	go build payments

start:
	./payments 

test-ping:
	curl localhost:3000/ping

test-payments:
	curl localhost:3000/payments

docker-build: 
	docker build -t payments .
	docker images

docker-run:
	docker run --name payments -td -p 3000:3000 payments
	docker ps

docker-network:
	docker network inspect host

docker-shell:
	docker exec -it payments bash 

docker-clean:
	docker stop payments
	docker rm payments
	docker rmi payments

docker-ip:
	docker-machine ip

heroku-tag:
	docker build -t payments .
	docker tag payments registry.heroku.com/pnguyen-goapi/web

heroku-push:
	heroku login 
	heroku container:login
	docker push registry.heroku.com/pnguyen-goapi/web
