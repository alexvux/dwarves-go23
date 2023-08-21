run:
	go run ./cmd/main.go
docker-build:
	docker build -t ex8 .
docker-run:
	docker run -p 8080:8080 ex8
docker-push:
	docker push ex8
docker-all: docker-build docker-push docker-run

.PHONY: run, docker-build, docker-run, docker-push, docker-all