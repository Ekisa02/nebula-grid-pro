# Nebula Grid

Nebula Grid is a modular Go project that demonstrates core primitives for distributed cloud systems: a registry for nodes and services, a simple leader-election (consensus) simulation, REST API using Gin, and persistent storage via Gorm.

This repository is intended as a reference architecture and starting point for building distributed control-plane components.

## Architecture

- cmd/cloudmesh: application entrypoint.
- internal/config: environment-backed configuration loader.
- internal/db: Gorm database initialization and migration helpers.
- internal/models: Gorm models (Node, Service).
- internal/api: HTTP handlers and small REST surface.
- internal/cluster: a simple leader-election simulation (for demo testing of distributed behaviors).
- internal/server: Gin router bootstrap and HTTP server initialization.
- pkg/logger: centralized logger setup.

Data flow summary:
- The HTTP API accepts node/service registrations and queries.
- The DB persists registry state (Gorm auto-migrate is enabled).
- A background consensus routine periodically reconciles leader status across registered nodes (simulated algorithm).

## Features

- Node and service registry with REST API
- Gorm-based persistence with support for sqlite (default) and PostgreSQL
- A simple consensus worker that elects a leader among known nodes
- Unit tests using an in-memory sqlite DB
- Dockerfile + docker-compose for quick local environment

## Installation

Prerequisites:
- Go 1.20+
- (Optional) Docker & docker-compose

Local build

1. Clone repository

   git clone https://github.com/yourusername/nebula-grid.git
   cd nebula-grid

2. Build

   make build

3. Run (defaults to sqlite in ./nebula.db)

   make run

With Docker Compose (Postgres)

1. Copy .env.example to .env and update if desired
2. Start services

   docker-compose up --build

## Configuration

Configuration is read from environment variables. Useful ones:
- PORT (default: 8080)
- DB_DRIVER (sqlite|postgres, default: sqlite)
- POSTGRES_DSN (postgres connection string if DB_DRIVER=postgres)
- CONSENSUS_INTERVAL (leader election tick in seconds, default: 10)

See .env.example for a starting point.

## API

- POST /api/v1/nodes  - register a node
- GET  /api/v1/nodes  - list nodes
- POST /api/v1/services - register a service
- GET  /api/v1/services - list services

Example: register a node

curl -X POST http://localhost:8080/api/v1/nodes -H 'Content-Type: application/json' -d '{"name":"node-a","address":"10.0.0.1:9000"}'

## Tests

Run unit tests with:

   go test ./...

## Notes on extending

- Replace the consensus simulation in internal/cluster with a real consensus/coordination algorithm (Raft, etcd) for production-grade behavior.
- Add authentication/authorization to the API.
- Add metrics and observability (Prometheus, OpenTelemetry).

## License

Apache-2.0
