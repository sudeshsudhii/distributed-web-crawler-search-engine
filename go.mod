module github.com/sudhi/distributed-search-engine

go 1.22.0

// Dependencies will be added as services are implemented.
// Core dependencies planned:
//   - google.golang.org/grpc (gRPC)
//   - github.com/redis/go-redis/v9 (Redis client)
//   - github.com/segmentio/kafka-go (Kafka client)
//   - github.com/jackc/pgx/v5 (PostgreSQL)
//   - github.com/prometheus/client_golang (Metrics)
//   - go.opentelemetry.io/otel (Tracing)
//   - github.com/golang-jwt/jwt/v5 (JWT auth)
//   - go.uber.org/zap (Structured logging)

require (
	github.com/redis/go-redis/v9 v9.5.1
	github.com/spaolacci/murmur3 v1.1.0
)
