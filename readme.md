go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

```bash
# if pgx should be the driver..makes sense if migrate is used as a library
go install -tags 'pgx' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# if postgres should be the driver..makes sense if migrate is used as a cli
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

migrate -source file://internal/db/migrations -database "$DATABASE_URL" up

migrate create --dir internal/db/migrations --ext sql create_products

if you get

when running migrate -> error: database driver: unknown driver postgresql/pgx (forgotten import?)

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```