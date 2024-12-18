# Go Project

## Requirements
- Go version **1.22.0** or higher.

To check your current Go version, run:
```bash
go version
```

To copy your current .env.example file and past it into .env file please run 
```bash
cp .env.example .env
```

Live server
```bash
air
```

To create a table (local db example)
```bash
goose -dir cmd/migrate/migrations create add_posts_table sql
```

To migrate the tables(up/down) (local db example)
```bash
goose -dir ./cmd/migrate/migrations postgres "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable" up/down
```
