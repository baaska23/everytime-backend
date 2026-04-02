# Everytime — Backend

University community platform for Mongolian students. `.edu` email auth only. Text-only content (no user image uploads). All content scoped to the user's university.

## Features
- **Board** — anonymous posts by faculty/department, upvote/downvote/report
- **Marketplace** — text listings, contact reveal on interest (no chat)
- **Lost & Found** — same contact reveal mechanic as marketplace
- **Courses** — course reviews, professor ratings
- **Timetable** — schedule builder
- **GPA Calculator** — client-side only, nothing stored
- **Ads** — top banner (only image maybe URL developed later), managed by admin
- **Admin** — moderation queue, account bans, ad slot management

## Stack
- Go 1.24+, PostgreSQL/Supabase, REST, Docker
- Vertical slice architecture (`internal/<feature>/handler|service|repository|models`)
- GORM for ORM

## Structure
```
myapp/
├── cmd/server/main.go
├── internal/
│   ├── shared/          # DB pool, logger, middleware, JWT helper
│   ├── auth/            # .edu email verify, login, register
│   ├── board/           # anonymous posts, comments, upvote/downvote, report
│   ├── marketplace/     # sell/buy listings, contact reveal, send offer
│   ├── courses/         # reviews, professor ratings
│   ├── timetable/       # schedule builder, conflict detection
│   ├── lostfound/       # lost & found listings (not for early stage of development)
│   └── admin/           # moderation dashboard, ban management, ad management
└── pkg/
    ├── apierror/        # standard HTTP error shapes
    ├── validator/       # IsEmail, IsEduEmail, IsNotEmpty
    ├── pagination/      # page/limit helpers
    └── ptr/             # pointer helpers
```

## Commands
```bash
go build ./...
go test ./...
go vet ./...
go mod tidy
```

## Conventions
- Wrap errors: `fmt.Errorf("context: %w", err)`
- No `init()`, no global state — inject via constructors
- Context as first param, propagated everywhere
- Parameterized queries only — never string formatting

## Agents
| Invoke | Does |
|--------|------|
| `use planner` | Implementation plan before building |
| `use architect` | Schema, service boundaries, infra |
| `use code-reviewer` | Quality, security, maintainability |
| `use go-reviewer` | Go idioms, concurrency, error handling |

## Skills (auto-loaded)
`golang` · `golang-patterns` · `golang-testing` · `api-design` · `backend-patterns` · `docker-patterns`