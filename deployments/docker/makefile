# Valores padrão para ambiente local
ENV ?= local

# Declara que estes comandos não são arquivos
.PHONY: db-local flyway local dev prd up-local help

# Comando para iniciar o banco local
db-local:
	docker compose up -d postgres-note

flyway:
	@if [ ! -f .env.$(ENV) ]; then \
		echo "Arquivo .env.$(ENV) não encontrado."; \
		exit 1; \
	fi
	@echo "Executando Flyway no ambiente: $(ENV)"
	@if [ "$(ENV)" = "local" ]; then \
		set -a; . ./.env.$(ENV); set +a; \
		export $$(cat .env.$(ENV) | grep -v '^#' | xargs) && \
		docker compose up flyway-note; \
	else \
	    set -a; . ./.env.$(ENV); set +a; \
       SQL_PATH="$$(pwd)/../sql_scripts/$(ENV)" && \
       echo "Using SQL path: $$SQL_PATH" && \
       export FLYWAY_URL=$$FLYWAY_URL && \
       export FLYWAY_USER=$$FLYWAY_USER && \
       export FLYWAY_PASSWORD=$$FLYWAY_PASSWORD && \
       export FLYWAY_LOCATIONS="filesystem:$$SQL_PATH" && \
       export FLYWAY_SCHEMAS=$$FLYWAY_SCHEMAS && \
       export FLYWAY_DATABASE=$$FLYWAY_DATABASE && \
       export FLYWAY_LOCK_RETRY_COUNT=100 && \
       export FLYWAY_MIXED=true && \
#        flyway repair && \
       flyway -X -baselineOnMigrate=true -outOfOrder=true -validateOnMigrate=false migrate; \
    fi

local:
	@$(MAKE) flyway ENV=local

dev:
	@$(MAKE) flyway ENV=dev

prd:
	@$(MAKE) flyway ENV=prd

# Comando para subir a aplicação completa no local
up-local: db-local local
	docker-compose up -d app-note

# Ajuda
help:
	@echo "Uso: make [comando]"
	@echo ""
	@echo "Comandos:"
	@echo "  local       Executa Flyway no ambiente local (via docker-compose)"
	@echo "  dev         Executa Flyway no ambiente de desenvolvimento (via docker run)"
	@echo "  prd         Executa Flyway no ambiente de produção (via docker run)"
	@echo "  flyway      Executa Flyway no ambiente especificado por ENV=ambiente"
	@echo "  db-local    Inicia apenas o banco de dados PostgreSQL local"
	@echo "  up-local    Inicia toda a pilha de aplicação local (DB + Flyway + App)"
	@echo "  help        Mostra esta ajuda"