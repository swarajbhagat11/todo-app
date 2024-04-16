test:
	go test ./...

build:
	docker-compose build

run:
	make build && docker-compose up -d && docker-compose logs -f

down:
	docker-compose down