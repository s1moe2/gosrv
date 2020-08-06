# gosrv

A golang RESTful API server with:
- nice decoupling (I hope) with a sort of repository pattern approach
- unit tests on the route handlers
- gorilla/mux for router
- logging middleware
- database migrations
- dockerfile

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