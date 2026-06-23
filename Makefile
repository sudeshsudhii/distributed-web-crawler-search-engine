# ============================================================================
# Distributed Web Crawler & Intelligent Search Engine — Build Automation
# ============================================================================

.PHONY: help build test lint fmt clean docker-build docker-up docker-down \
        bench proto migrate seed

# Default target
.DEFAULT_GOAL := help

# ============================================================================
# Variables
# ============================================================================

GO := go
GOFLAGS := -v
BINARY_DIR := bin
SERVICES := crawler-coordinator url-frontier crawler-worker search-api \
            ranking-service indexing-service auth-service analytics-service

DOCKER_COMPOSE := docker compose
DOCKER_COMPOSE_FILE := docker-compose.yml

# Version info (injected at build time)
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u '+%Y-%m-%dT%H:%M:%SZ')
LDFLAGS := -ldflags "-X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME)"

# ============================================================================
# Help
# ============================================================================

help: ## Show this help message
	@echo ""
	@echo "  Distributed Search Engine — Build Targets"
	@echo "  =========================================="
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""

# ============================================================================
# Build
# ============================================================================

build: ## Build all Go services
	@echo "==> Building all services..."
	@mkdir -p $(BINARY_DIR)
	@for svc in $(SERVICES); do \
		echo "  Building $$svc..."; \
		$(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BINARY_DIR)/$$svc ./cmd/$$svc/; \
	done
	@echo "==> All services built successfully."

build-%: ## Build a specific service (e.g., make build-search-api)
	@echo "==> Building $*..."
	@mkdir -p $(BINARY_DIR)
	$(GO) build $(GOFLAGS) $(LDFLAGS) -o $(BINARY_DIR)/$* ./cmd/$*/
	@echo "==> $* built successfully."

# ============================================================================
# Test
# ============================================================================

test: ## Run all Go tests with coverage
	@echo "==> Running tests..."
	$(GO) test -race -coverprofile=coverage.out ./...
	$(GO) tool cover -func=coverage.out | tail -n 1
	@echo "==> Tests complete."

test-verbose: ## Run tests with verbose output
	$(GO) test -race -v -coverprofile=coverage.out ./...

test-coverage: test ## Generate HTML coverage report
	$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "==> Coverage report: coverage.html"

test-integration: ## Run integration tests
	$(GO) test -race -tags=integration -v ./tests/integration/...

# ============================================================================
# Code Quality
# ============================================================================

lint: ## Run linters (requires golangci-lint)
	@echo "==> Running linters..."
	golangci-lint run ./...

fmt: ## Format Go code
	@echo "==> Formatting code..."
	$(GO) fmt ./...
	goimports -w .

vet: ## Run go vet
	$(GO) vet ./...

# ============================================================================
# Docker
# ============================================================================

docker-build: ## Build all Docker images
	@echo "==> Building Docker images..."
	@for svc in $(SERVICES); do \
		echo "  Building image: $$svc"; \
		docker build -t distributed-search/$$svc:$(VERSION) -f cmd/$$svc/Dockerfile .; \
	done
	@echo "==> Building Python services..."
	docker build -t distributed-search/parser:$(VERSION) -f services/parser/Dockerfile .
	docker build -t distributed-search/ai-features:$(VERSION) -f services/ai-features/Dockerfile .

docker-up: ## Start local dev infrastructure (PostgreSQL, Redis, Kafka, etc.)
	@echo "==> Starting infrastructure..."
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) up -d
	@echo "==> Infrastructure is running."
	@echo "  PostgreSQL:  localhost:5432"
	@echo "  Redis:       localhost:6379"
	@echo "  Kafka:       localhost:9092"
	@echo "  MinIO:       localhost:9000 (console: 9001)"
	@echo "  Prometheus:  localhost:9090"
	@echo "  Grafana:     localhost:3000 (admin/admin)"

docker-down: ## Stop local dev infrastructure
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down

docker-clean: ## Stop and remove all containers, volumes, and images
	$(DOCKER_COMPOSE) -f $(DOCKER_COMPOSE_FILE) down -v --rmi local

# ============================================================================
# Database
# ============================================================================

migrate-up: ## Run database migrations
	@echo "==> Running migrations..."
	migrate -path db/migrations -database "postgres://crawler:crawler@localhost:5432/search_engine?sslmode=disable" up

migrate-down: ## Rollback database migrations
	migrate -path db/migrations -database "postgres://crawler:crawler@localhost:5432/search_engine?sslmode=disable" down 1

seed: ## Seed the database with test data
	@echo "==> Seeding database..."
	$(GO) run ./scripts/seed-data/main.go

# ============================================================================
# Benchmarks
# ============================================================================

bench: ## Run Go benchmarks
	@echo "==> Running benchmarks..."
	$(GO) test -bench=. -benchmem ./internal/...

bench-search: ## Benchmark search latency
	$(GO) test -bench=BenchmarkSearch -benchmem -count=5 ./internal/search/...

bench-index: ## Benchmark indexing throughput
	$(GO) test -bench=BenchmarkIndex -benchmem -count=5 ./internal/index/...

# ============================================================================
# Cleanup
# ============================================================================

clean: ## Remove build artifacts
	@echo "==> Cleaning..."
	rm -rf $(BINARY_DIR)
	rm -f coverage.out coverage.html
	$(GO) clean -cache -testcache
	@echo "==> Clean complete."
