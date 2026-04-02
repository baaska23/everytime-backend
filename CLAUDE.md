# everytime-backend

Go microservice backend using clean architecture.

## Stack

- Go 1.25+, PostgreSQL, gRPC + REST (grpc-gateway), Docker
- sqlc for type-safe SQL, Wire for dependency injection
- Clean architecture: domain → repository → service → handler

## Commands

```bash
go build ./...          # Build
go test ./...           # Test
go vet ./...            # Lint
go mod tidy             # Clean dependencies
```

## Agents

Use these by name when asking Claude for help:

| Command | Agent | What it does |
|---------|-------|-------------|
| "use planner" | `planner` | Creates implementation plans for features — run before building anything complex |
| "use architect" | `architect` | Designs schemas, service boundaries, infrastructure decisions |
| "use code-reviewer" | `code-reviewer` | General code review for quality, security, maintainability |
| "use go-reviewer" | `go-reviewer` | Go-specific review: idioms, concurrency, error handling, performance |

## Skills (auto-loaded reference knowledge)

Skills are in `skills/` and get loaded automatically when relevant. You can also invoke them with `/skill-name`:

| Skill | When it activates |
|-------|------------------|
| `golang` | Writing or editing any `.go` file |
| `golang-patterns` | Go idioms, package design, interfaces |
| `golang-testing` | Writing tests, benchmarks, fuzz tests, TDD |
| `api-design` | Designing REST endpoints, pagination, errors |
| `backend-patterns` | Repository/service layers, caching, middleware |
| `docker-patterns` | Dockerfile, docker-compose, container setup |

## Conventions

- Follow Effective Go and Go Code Review Comments
- Wrap errors with `fmt.Errorf("context: %w", err)` — never string match
- No `init()` functions — explicit initialization only
- No global mutable state — pass dependencies via constructors
- Context as first parameter, propagated through all layers
- Use GORM
- Migrations in `migrations/` — never alter DB directly
- Parameterized queries only (`$1`, `$2`) — never string formatting
