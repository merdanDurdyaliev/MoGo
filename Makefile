build:
	docker-compose build

run:
	docker-compose up -d --build

test:
	go test -v ./...

swag:
	swag init -g cmd/main.go

