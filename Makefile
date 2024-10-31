SHELL := /bin/bash
GO := docker run --rm -it -v $(PWD):/usr/src/myapp -v $(GOPATH):/go -w /usr/src/myapp -p 8080:8080 golang:1.23.0 go

.PHONY: up
up:
	@echo "--> Create network"
	docker network create test_project_golang

	@echo "--> Docker Compose Up..."
	docker-compose -f deployments/docker-compose.yml up -d --build

.PHONY: down
down:
	@echo "--> Docker Compose Down..."
	docker-compose -f deployments/docker-compose.yml down
	- docker network rm test_project_golang

.PHONY: restart
restart: down up
