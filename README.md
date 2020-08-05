# gosrv

A golang RESTful API server with:
- nice decoupling (I hope) with a sort of reposiroty pattern approach
- unit tests on the route handlers
- database migrations
- dockerfile
- gorilla/mux for router

### Requirements

Create a PostgreSQL database and user:
```postgresql
create database thedb;
create user theuser with encrypted password 'thepwd';
grant all privileges on database thedb to theuser;
```

### Tests

Unit tests are kept alongside their respective source files.
Run tests with `make test`.

### Migrations

- Create a migration
    > make migrate-create MIGRATION_NAME="test"

- Run migrations up/down
    > make migrate-[up/down] DB_URI="postgres://dbuser:dbpwd@statuspsql:5432/dbname?sslmode=disable"