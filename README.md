# test_case_putri

## Setup migration

### Install goose

```bash
cd migration
go install github.com/pressly/goose/v3/cmd/goose@latest
```
### Create file migration

```bash
goose create create_new_table sql
```
### Up file migration
```bash
goose mysql 'user:root@/dbname?parseTime=true' up
```

### Down file migration
```bash
goose mysql 'user:root@/dbname?parseTime=true' down
```

## How to run

### Local

```bash
go run main.go
```
