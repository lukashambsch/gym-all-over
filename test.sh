# shell script to set up env and run tests
export GOENV=test
go get -u github.com/golang/dep/...
psql -c "CREATE USER $POSTGRES_ENV_POSTGRES_USER WITH PASSWORD '$POSTGRES_ENV_POSTGRES_PASSWORD';"
echo "CREATE EXTENSION IF NOT EXISTS pgcrypto" | psql -d postgres
go run store/migrations/migrate.go -up
go test -cover -v $(go list ./... | grep -v /vendor/)
go run store/migrations/migrate.go -down
export GOENV=local
