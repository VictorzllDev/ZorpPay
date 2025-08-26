# Configurações
APP_NAME      := ZorpPay
BINARY_DIR    := bin
SWAGGER_DIR   := docs
MAIN_PATH     := cmd/api/main.go
VERSION       := 0.5.0
GOFLAGS       := -ldflags="-s -w -X main.version=$(VERSION)"

# Cores (compatíveis com Kitty)
RED    := $$(printf '\033[31m')
GREEN  := $$(printf '\033[32m')
YELLOW := $$(printf '\033[33m')
RESET  := $$(printf '\033[0m')

.PHONY: dev run build swagger deps test lint clean help

## dev: Inicia a API em modo desenvolvimento (com Swagger)
dev: swagger
	@printf "${GREEN}🌍 Iniciando API em modo desenvolvimento...${RESET}\n"
	@go run $(GOFLAGS) $(MAIN_PATH)

## run: Inicia a API sem gerar documentação
run:
	@printf "${GREEN}🌍 Iniciando API...${RESET}\n"
	@go run $(MAIN_PATH)

## build: Compila o binário da API
build:
	@printf "${GREEN}🔨 Compilando ${APP_NAME} v${VERSION}...${RESET}\n"
	@go build $(GOFLAGS) -o $(BINARY_DIR)/$(APP_NAME) $(MAIN_PATH)

## swagger: Gera/atualiza documentação Swagger
swagger:
	@if ! command -v swag &> /dev/null; then \
		printf "${RED}❌ 'swag' não instalado.${RESET}"; \
		exit 1; \
	fi
	@printf "${GREEN}📝 Gerando documentação Swagger...${RESET}\n"
	@swag init -g $(MAIN_PATH) -o $(SWAGGER_DIR)
	@printf "${GREEN}✅ Documentação gerada em: ${SWAGGER_DIR}/${RESET}\n"

## deps: Instala/atualiza dependências
deps:
	@printf "${GREEN}📦 Gerenciando dependências...${RESET}\n"
	@go mod tidy
	@go mod download

## lint: Executa análise estática do código
lint:
	@printf "${GREEN}🔍 Verificando qualidade do código...${RESET}\n"
	@if ! command -v golangci-lint &> /dev/null; then \
		printf "${RED}❌ 'golangci-lint' não instalado.${RESET}"; \
		exit 1; \
	fi
	@golangci-lint run ./...

## clean: Remove arquivos gerados
clean:
	@printf "${GREEN}🧹 Limpando artefatos...${RESET}\n"
	@rm -rf $(BINARY_DIR) $(SWAGGER_DIR)
	@go clean

## help: Mostra esta ajuda
help:
	@printf "${YELLOW}\n📖 Comandos disponíveis para ${APP_NAME}:${RESET}\n"
	@printf "${GREEN}dev${RESET}: Inicia a API em modo desenvolvimento (com Swagger)\n"
	@printf "${GREEN}run${RESET}: Inicia a API sem gerar documentação\n"
	@printf "${GREEN}build${RESET}: Compila o binário da API\n"
	@printf "${GREEN}swagger${RESET}: Gera/atualiza documentação Swagger\n"
	@printf "${GREEN}deps${RESET}: Instala/atualiza dependências\n"
	@printf "${GREEN}lint${RESET}: Executa análise estática do código\n"
	@printf "${GREEN}clean${RESET}: Remove arquivos gerados\n"
	@printf "${GREEN}help${RESET}: Mostra esta ajuda\n"

.DEFAULT_GOAL := help
