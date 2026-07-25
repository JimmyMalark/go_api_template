# Go API Template

A production-ready REST API template built with modern Go practices.

This template is designed to be cloned whenever you start a new backend project, so you don't have to rebuild the same infrastructure every time.

## Features

- 🚀 Go 1.26+
- 🌐 Chi router
- 🐘 PostgreSQL
- ⚡ sqlc for type-safe SQL
- 🔄 Database migrations
- 📦 Redis caching
- 📄 OpenAPI (Swagger)
- ✅ Request validation
- 📄 Structured API responses
- 📚 Pagination support
- 🪵 Structured logging
- ⚙️ Environment-based configuration
- 🐳 Docker & Docker Compose
- 🛠 Taskfile automation
- 🔒 Layered architecture

# Getting Started

## 1. Clone the template

```bash
git clone <repository-url> my-project
cd my-project
```

Rename the Go module:

```bash
task rename MODULE=github.com/yourname/my-awesome-api
```

---

## 2. Configure the application

Copy the example environment file.

```bash
cp .env.example .env
```

Update the values inside `.env`.

---

## 3. Start dependencies

```bash
task docker:up
```

---

## 4. Run migrations

```bash
task db:migrate
```

---

## 5. Generate sqlc code

```bash
task db:sqlc:generate
```

---

## 6. Start the API

```bash
task dev
```

or

```bash
task run
```

The API will now be available at

```
http://localhost:8080
```

---

# Configuration

Configuration is loaded from environment variables.

See `.env.example` for all available options.

---

# API Documentation

Swagger/OpenAPI documentation is available at

```
/docs
```

---

# Architecture

The project follows a layered architecture.

```text
HTTP
    │
Handlers
    │
Services
    │
Repositories
    │
PostgreSQL
```

Caching is handled inside the service layer.

Repositories are responsible only for database access.

---

# Technologies

- Go
- Chi
- PostgreSQL
- sqlc
- Redis
- Docker
- Taskfile
- Cleanenv
- OpenAPI / Swagger

---

# Development

Useful commands:

```bash
task dev          # Run with hot reload
task run          # Run normally
task test         # Run tests
task lint         # Run linter
task fmt          # Format code
task db:generate  # Generate sqlc code
task db:migrate   # Run migrations
task docker:up    # Start containers
task docker:down  # Stop containers
```

---

# Planned Features

- Authentication
- Authorization
- Refresh tokens
- Background jobs
- Email support
- File uploads
- Metrics
- OpenTelemetry
- Integration testing
- CI/CD examples

---

# License

MIT