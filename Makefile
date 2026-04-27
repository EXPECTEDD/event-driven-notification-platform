include .env
export

POSTGRES_HOST := $(or $(POSTGRES_HOST),localhost)
POSTGRES_PORT := $(or $(POSTGRES_PORT),5432)
POSTGRES_SSL := $(or $(POSTGRES_SSL),disable)

export PROJECT_ROOT=$(shell pwd)
export UID=$(shell id -u)
export GID=$(shell id -g)

postgres-up:
	@mkdir -p $(PROJECT_ROOT)/out/pgdata; \
	$(DOCKER_COMPOSE) up -d postgres;

postgres-down:
	@$(DOCKER_COMPOSE) down postgres;

postgres-cleanup:
	@read -p "Очистить все данные postgres? Опасность утери данных. [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		$(DOCKER_COMPOSE) down postgres && \
		sudo rm -rf $(PROJECT_ROOT)/out/pgdata && \
		echo "Данные удалены."; \
	else \
		echo "Отмена удаления данных."; \
	fi;

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Отсутствует необходимый параметр seq. Пример: make migrate-create seq=init"; \
	fi; \
	mkdir -p $(PROJECT_ROOT)/migrations; \
	$(DOCKER_COMPOSE) run --rm --user $(UID):$(GID) migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq $(seq);

migrate-up:
	@$(DOCKER_COMPOSE) run --rm migrate \
		-path /migrations \
		-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSL) \
		up $(n);

migrate-down:
	@$(DOCKER_COMPOSE) run --rm migrate \
		-path /migrations \
		-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=$(POSTGRES_SSL) \
		down $(n);

port-forwarder-up:
	@$(DOCKER_COMPOSE) up -d port-forwarder;

port-forwarder-down:
	@$(DOCKER_COMPOSE) down port-forwarder;