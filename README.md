# SOCIAL-GO

SOCIAL-GO is a small social media backend implemented in Go. It provides an HTTP API, database migrations, and simple in-memory/file-backed storage helpers intended for local development and learning.

## Repository layout

- `cmd/` - application entrypoints
  - `api/` - HTTP API server (main, handlers)
  - `migrate/` - database migration runner and SQL migration files
- `internal/` - internal packages
  - `env/` - environment helpers
  - `store/` - data storage interfaces and implementations for users and posts
  - `db/` - database helpers
- `bin/` - compiled binaries (committed for convenience)
- `scripts/` - auxiliary scripts (DB init, etc.)
- `migrate/` - migration tooling
- `.env` - environment variables for local development

## Requirements

- Go 1.20+ (or the version used in the project)
- PostgreSQL (for running migrations and the app in production-like mode)
- make (optional) and Docker Compose if you want to run services in containers

## Quick start (local)

1. Copy `.env` and adjust values as needed (or create a local override):

```bash
cp .env .env.local
# edit .env.local
```

2. Run database (example using docker-compose):

```bash
docker compose up -d postgres
```

3. Run migrations:

```bash
make migrate-up
```

4. Build and run the API server:

```bash
go build -o bin/main ./cmd/api
./bin/main
```

Or with Docker Compose (if defined):

```bash
docker compose up --build
```

## Development notes

- Migrations live under `cmd/migrate/migrations/` and are plain SQL files named with sequence prefixes.
- Storage implementations are under `internal/store` and can be extended to use a real DB layer.
- The project intentionally includes compiled binaries in `bin/` for convenience; you may want to add `bin/` and `.env` to `.gitignore` for future work.

## Tests

Run `go test ./...` from the project root.

## Contributing

Open an issue or PR. Keep commits small and focused. The repo owner/maintainer is `aryan735`.

## License

This repository does not include a LICENSE file. Add one if you plan to make the project public.
