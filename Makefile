.PHONY: gen
req:
	go get entgo.io/ent/cmd/ent
up:
	docker-compose up -d

db-reset:
	docker-compose rm --stop --force db
	docker-compose up -d db

run:
	go run ./cmd/badge-issue-service

test:
	go test --cover ./...

build:
	go build -o web3-identity-provider cmd/main.go

docker-image:
	DOCKER_BUILDKIT=1 docker build -t web3-identity-provider:latest .
