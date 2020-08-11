# gosrv

A golang RESTful API server with:
- dependency injection on handlers and a sort of repository pattern approach for data layer
- unit tests on the route handlers
- gorilla/mux for router
- request logging middleware
- OpenAPI documentation
- SwaggerUI to serve API docs
- database migrations
- dockerfile

### Requirements

Create a PostgreSQL database and user:
```postgresql
CREATE DATABASE thedb;
CREATE USER theuser WITH ENCRYPTED PASSWORD 'thepwd';
GRANT ALL PRIVILEGES ON DATABASE thedb TO theuser;
```

### Tests

Unit tests are kept alongside their respective source files.
Run tests with `make test`.

### Migrations

- Create a migration
    > make migrate-create MIGRATION_NAME="test"

- Run migrations up/down
    > make migrate-[up/down] DB_URI="postgres://dbuser:dbpwd@statuspsql:5432/dbname?sslmode=disable"