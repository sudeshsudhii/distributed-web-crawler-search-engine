# 🔍 Distributed Web Crawler & Intelligent Search Engine

[![CI](https://github.com/sudhi/distributed-search-engine/actions/workflows/ci.yml/badge.svg)](https://github.com/sudhi/distributed-search-engine/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sudhi/distributed-search-engine)](https://goreportcard.com/report/github.com/sudhi/distributed-search-engine)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A **production-grade distributed search system** built from scratch — no Elasticsearch, no Solr. Demonstrates mastery of distributed systems, information retrieval, and modern search infrastructure.

> **M.Tech Major Project** — Designed to be significantly more advanced than typical student projects, suitable for FAANG/product company placements and research publication.

---

## ✨ Key Features

| Category | Feature | Status |
|:---|:---|:---:|
| 🕷️ **Crawler** | Distributed crawling at 1000+ pages/sec across 100 workers | 🔲 |
| 🕷️ **Crawler** | robots.txt compliance, politeness policies, domain sharding | 🔲 |
| 🕷️ **Crawler** | Bloom filter URL deduplication, SimHash content deduplication | 🔲 |
| 🕷️ **Crawler** | Adaptive crawl scheduling with reinforcement learning (DQN) | 🔲 |
| 📊 **Index** | Custom compressed inverted index (VByte encoding, 10x compression) | 🔲 |
| 📊 **Index** | Distributed sharding with configurable allocation strategies | 🔲 |
| 🔍 **Search** | BM25 ranking with multi-shard scatter-gather | 🔲 |
| 🔍 **Search** | PageRank & HITS graph-based ranking algorithms | 🔲 |
| 🧠 **AI** | Semantic vector search (sentence-transformers + HNSW) | 🔲 |
| 🧠 **AI** | Hybrid retrieval (BM25 + Vector) with Reciprocal Rank Fusion | 🔲 |
| 🧠 **AI** | Cross-encoder neural reranking for top-k precision | 🔲 |
| ☸️ **Infra** | Kubernetes-native with HPA auto-scaling, Helm charts | 🔲 |
| 📡 **Observability** | Prometheus + Grafana + OpenTelemetry + Loki full stack | 🔲 |
| 🔒 **Security** | JWT (RS256) auth, RBAC, rate limiting, mTLS | 🔲 |

---

## 🏗️ Architecture

```
14 Microservices | Go + Python | Kafka | Redis | PostgreSQL | Kubernetes
```

### System Overview

```
┌─────────────────────────────────────────────────────────────────┐
│                        API Gateway (Kong)                       │
│                    Rate Limiting + Auth (JWT)                    │
├─────────────────────────┬───────────────────────────────────────┤
│    Search Pipeline      │          Crawl Pipeline               │
│  ┌──────────────────┐   │  ┌───────────────────────────────┐   │
│  │   Search API     │   │  │   Crawler Coordinator (Raft)  │   │
│  │   Ranking Svc    │   │  │   URL Frontier (Redis)        │   │
│  │   AI Features    │   │  │   Worker Cluster (N pods)     │   │
│  └──────────────────┘   │  │   Parser + Dedup Service      │   │
│                         │  └───────────────────────────────┘   │
├─────────────────────────┴───────────────────────────────────────┤
│                    Kafka Event Bus                              │
├─────────────────────────────────────────────────────────────────┤
│  PostgreSQL  │  Redis Cluster  │  MinIO (S3)  │  Vector Store  │
├─────────────────────────────────────────────────────────────────┤
│     Prometheus  │  Grafana  │  OpenTelemetry  │  Loki          │
└─────────────────────────────────────────────────────────────────┘
```

---

## 📊 Benchmark Targets

| Metric | Target | Status |
|:---|:---|:---:|
| Crawl Throughput | > 1,000 pages/sec (cluster) | 🔲 |
| Search Latency (p50) | < 50ms | 🔲 |
| Search Latency (p99) | < 200ms | 🔲 |
| Index Compression Ratio | > 10x | 🔲 |
| Search QPS | > 1,000 queries/sec | 🔲 |
| Documents Indexed | 1,000,000+ | 🔲 |

---

## 🛠️ Technology Stack

| Layer | Technology | Why |
|:---|:---|:---|
| **Core Services** | Go 1.22+ | Best-in-class concurrency, compiled performance |
| **ML/AI Services** | Python 3.12+ | ML ecosystem, ONNX Runtime |
| **Message Queue** | Apache Kafka (KRaft) | High throughput, replayability |
| **Cache / Frontier** | Redis 7+ Cluster | Sub-ms latency, sorted sets |
| **Metadata DB** | PostgreSQL 16+ | ACID, rich indexing |
| **Object Storage** | MinIO | S3-compatible, self-hosted |
| **Orchestration** | Kubernetes 1.30+ | Auto-scaling, self-healing |
| **Monitoring** | Prometheus + Grafana | Industry standard |
| **Tracing** | OpenTelemetry + Tempo | Vendor-neutral distributed tracing |
| **Logging** | Loki + Grafana Alloy | Efficient log aggregation |

---

## 📁 Project Structure

```
distributed-search-engine/
├── cmd/                    # Service entry points (Go)
│   ├── crawler-coordinator/
│   ├── url-frontier/
│   ├── crawler-worker/
│   ├── search-api/
│   ├── ranking-service/
│   ├── indexing-service/
│   ├── auth-service/
│   └── analytics-service/
├── internal/               # Shared Go packages
│   ├── crawler/            # HTTP fetching, robots.txt
│   ├── frontier/           # Priority queue, Bloom filter
│   ├── index/              # Inverted index, compression
│   ├── search/             # BM25, PageRank, hybrid
│   ├── dedup/              # SimHash, MinHash
│   ├── auth/               # JWT, RBAC
│   └── common/             # Config, logging, metrics
├── services/               # Python services
│   ├── parser/             # HTML parsing, text extraction
│   └── ai-features/        # Embeddings, reranking
├── deployments/            # Kubernetes manifests
├── configs/                # Configuration files
├── tests/                  # Integration & load tests
├── docs/                   # Documentation
└── web/                    # Admin dashboard UI
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.22+
- Python 3.12+
- Docker & Docker Compose
- Make

### Quick Start

```bash
# Clone the repository
git clone https://github.com/sudhi/distributed-search-engine.git
cd distributed-search-engine

# Start infrastructure (PostgreSQL, Redis, Kafka, MinIO, etc.)
make docker-up

# Build all services
make build

# Run tests
make test

# Run benchmarks
make bench
```

### Kubernetes Deployment

```bash
# Build Docker images
make docker-build

# Deploy with Helm
helm install search-engine deployments/helm/search-engine/
```

---

## 📈 Implementation Roadmap

| Phase | Weeks | Focus | Status |
|:---|:---:|:---|:---:|
| 1. Foundation | 1–4 | Crawler, parser, frontier, basic pipeline | 🔲 |
| 2. Core Indexing | 5–8 | Inverted index, BM25, Search API | 🔲 |
| 3. Distribution | 9–12 | Kafka, multi-worker, sharding, leader election | 🔲 |
| 4. Advanced Ranking | 13–16 | PageRank, HITS, semantic search, hybrid retrieval | 🔲 |
| 5. Production Hardening | 17–20 | Auth, K8s, observability, dedup | 🔲 |
| 6. Polish & Research | 21–24 | Benchmarks, UI, RL adaptive crawling, thesis | 🔲 |

---

## 📚 Research Publications

This project explores publication-worthy extensions:

1. **RL-Based Adaptive Crawling** — DQN agent for intelligent crawl scheduling
2. **GNN-Based Search Ranking** — GraphSAGE over web graph for ranking
3. **Temporal-Aware PageRank** — Decay-weighted link analysis

---

## 📄 License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

---

## 🙏 Acknowledgments

Inspired by the architectures of Google Search, Apache Nutch, Elasticsearch, and Vespa. Built with insights from Google, Netflix, Uber, and LinkedIn engineering blogs.
