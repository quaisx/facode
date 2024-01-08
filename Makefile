install:
	go mod download

test:
	go clean --testcache
	go test ./...

run-server:
	go run cmd/server/main.go

run-client:
	go run cmd/client/main.go

run-docker:
	docker-compose up --abort-on-container-exit --force-recreate --build server --build client