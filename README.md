# cloud-simple-demo

`cloud-simple-demo` is an end-to-end application showing how the Golang Acexy cloud modules turn database metadata into a structured HTTP service. It demonstrates generated models, typed repositories, business services, standard CRUD routers, time-range pagination, and row-level authority control.

## What It Demonstrates

```text
HTTP request
  -> starter-gin Router
  -> cloud-web BaseRouter
  -> generated BizService
  -> cloud-database Repository
  -> starter-gorm Mapper
  -> MySQL
```

The project contains two related resources:

- `department` demonstrates standard generated CRUD behavior.
- `employee` demonstrates standard CRUD, separate user/admin routers, time-range pagination, and authority-controlled data access.

Generated files provide the repeatable application foundation and stay aligned with the current cloud module contracts.

## Requirements

- Go `1.26.7` or later
- MySQL reachable at `127.0.0.1:13306`
- Database `test`
- Demo tables matching the generated `Department` and `Employee` models

The credentials in `cmd/main.go` are local demonstration defaults. Replace them before using the project outside an isolated development environment.

## Main Components

- `cmd/main.go` assembles GORM and Gin through `starter-parent`.
- `internal/model/*_gen.go` contains generated database models and DTOs.
- `internal/service/repo` contains typed GORM mappers and cloud repositories.
- `internal/service/biz` implements `cloud-web` business-service contracts.
- `internal/handler/rest/usr` exposes user-facing authority-controlled routes.
- `internal/handler/rest/adm` exposes administrative CRUD routes.
- `scripts/verify.sh` exercises the complete HTTP flow.
- `test/test.http` provides request examples for compatible IDE HTTP clients.

## Application Startup

The application registers infrastructure in dependency order:

1. `GormStarter` opens the MySQL connection.
2. `GinStarter` registers generated routers.
3. `StarterLoader` starts both components and coordinates lifecycle state.

Run the application from this module:

```bash
go run ./cmd
```

The HTTP server listens on `:8080` by default.

## Generated Architecture

Each generated resource follows the same typed layers:

```text
Model and DTOs
  -> BaseMapper
  -> Repository
  -> BaseBizService
  -> BaseRouter
```

Repository methods use business verbs such as `Save`, `Query`, `Modify`, and `Remove`; the embedded GORM mapper uses SQL-oriented verbs such as `Insert`, `Select`, `Update`, and `Delete`.

Generated repositories expose `Columns()` together with `Wrapper()`,
`PageWrapper()`, and `ModifyWrapper()`. Custom business queries can therefore
compose type-safe conditions without constructing a Mapper directly:

```go
c := repo.Columns()
records, err := repo.QueryByWrapper(
	repo.Wrapper().
		Eq(c.Status, int32(1)).
		OrderByDesc(c.CreatedAt),
)
```

The generic REST bridge keeps conditions and updates as database-column Maps.
This preserves explicit zero values while generated BizServices translate the
request into `rds.NewMapQuery` or `rds.NewMapPageQuery` calls.

## Authority Control

The user-facing employee router is created with `NewBaseRouterWithAuthority`. The router resolves the current authority identity and controls the configured struct field and database column rather than trusting client input.

Authority-aware save, query, modify, and remove flows therefore apply the server-derived authority value consistently. Administrative routers can use a different authority provider or expose unrestricted operations according to application policy.

## Time-Range Pagination

Generated department and employee services allow pagination filters on `created_at` and `updated_at`. An omitted `field` uses `created_at`; explicit camel-case fields such as `updatedAt` are normalized before whitelist validation.

Each range requires at least one bound. When both bounds are present, the interval is left-closed and right-open: `[start, end)`.

```json
{
  "number": 1,
  "size": 10,
  "condition": {
    "name": "Alice"
  },
  "timeRanges": [
    {
      "field": "createdAt",
      "start": 1704067200000
    }
  ]
}
```

## Extending Generated Code

Keep generated foundations and custom behavior separate:

- Add model-independent business methods in the corresponding BizService extension file.
- Add specialized SQL or repository operations in a Repository or Mapper extension.
- Add non-standard HTTP endpoints on the concrete router.
- Override a standard HTTP handler by defining the same method on the concrete
  router; `RegisterBaseHandlers(router, r)` dispatches to that method and keeps
  embedded defaults for handlers that are not overridden.
- Preserve generated base contracts so the standard CRUD flow remains available.

## Verification

After MySQL and the application are running, inspect `scripts/verify.sh` before executing it, then use it to validate create, query, update, delete, time-range pagination, and authority-related behavior.

The demo intentionally depends on local infrastructure and is not a unit-test fixture. Adapt endpoints, credentials, and schema setup to the local environment.

## Design Notes

- This project demonstrates the complete cloud stack; it is not a reusable library.
- Database-generated IDs and timestamps are not supplied by clients.
- DTO definitions control writable and queryable fields exposed by base routers.
- Empty query Maps are valid, while empty update and delete conditions remain protected.
- Custom code should extend generated types instead of duplicating the framework's base CRUD pipeline.
