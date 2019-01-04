all: run

name = fesl-backend
namespace = localhost/go-heroes
entrypoint = ./cmd/$(name)
binary = $(name)

deps:
	$(call colorized, "Downloading dependencies via docker...")
	glide install
.PHONY: deps

codegen:
	$(call colorized, "Generating code...")
	go generate $(entrypoint)
.PHONY: codegen

api:
	$(MAKE) compile name=heroes-api
.PHONY: api

backend:
	$(MAKE) compile name=fesl-backend
.PHONY: backend

compile:
	$(call colorized, "Compiling Golang code as binary...")
	CGO_ENABLED=0 \
		go build -v -o $(binary) -ldflags='-w -s' $(entrypoint)
.PHONY: compile

run: compile start
.PHONY: run

up: compile
	docker-compose up --build backend

start:
	$(call colorized, "Starting compiled binary...")
	./$(binary) --config dev.env
.PHONY: start

clean:
	$(call colorized, "Removing compiled binary...")
	$(rm) $(binary)
.PHONY: clean

grant:
	chown $(shell stat -c '%u:%g' .) $(binary)
.PHONY: grant

define colorized
	@tput setaf 6
	@echo $1
	@tput sgr0
endef
