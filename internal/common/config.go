// Package common provides shared utilities used across all services including
// configuration management, structured logging, metrics registration, and
// common error types.
package common

// Config holds the global application configuration.
type Config struct {
	// Service identity
	ServiceName string
	Version     string
	Environment string // "dev", "staging", "production"

	// Server
	HTTPPort int
	GRPCPort int

	// PostgreSQL
	PostgresDSN string

	// Redis
	RedisAddr     string
	RedisPassword string

	// Kafka
	KafkaBrokers []string

	// MinIO
	MinIOEndpoint  string
	MinIOAccessKey string
	MinIOSecretKey string

	// Observability
	PrometheusPort int
	OTelEndpoint   string
	LogLevel       string // "debug", "info", "warn", "error"
}

// DefaultConfig returns a configuration with development defaults.
func DefaultConfig() *Config {
	return &Config{
		ServiceName:    "search-engine",
		Version:        "dev",
		Environment:    "dev",
		HTTPPort:       8080,
		GRPCPort:       9090,
		PostgresDSN:    "postgres://crawler:crawler@localhost:5432/search_engine?sslmode=disable",
		RedisAddr:      "localhost:6379",
		RedisPassword:  "",
		KafkaBrokers:   []string{"localhost:9094"},
		MinIOEndpoint:  "localhost:9000",
		MinIOAccessKey: "minioadmin",
		MinIOSecretKey: "minioadmin",
		PrometheusPort: 2112,
		OTelEndpoint:   "localhost:4317",
		LogLevel:       "info",
	}
}
