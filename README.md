<!-- Autogenerated by Typical-Go. DO NOT EDIT. -->

# Typical-RESTful-Server

Example of typical and scalable RESTful API Server for Go

## Getting Started

### Prerequisite

1. Install [Go](https://golang.org/doc/install) or `brew install go`

### Run

Use `./typicalw run` to compile and run local development. You can find the binary at `bin` folder

## Release Distribution

Use `./typicalw release` to make the release. You can find the binary at `release` folder. More information check [here](https://typical-go.github.io/release.html)

## Configuration

### Server

| Key | Type | Default | Required | Description |	
|---|---|---|---|---|
|SERVER_DEBUG|True or False|false|||

### Postgres Database

| Key | Type | Default | Required | Description |	
|---|---|---|---|---|
|PG_DBNAME|String||true||
|PG_USER|String|postgres|true||
|PG_PASSWORD|String|pgpass|true||
|PG_HOST|String|localhost|||
|PG_PORT|Integer|5432|||

### redis

| Key | Type | Default | Required | Description |	
|---|---|---|---|---|
|REDIS_HOST|String|localhost|true||
|REDIS_PORT|String|6379|true||
|REDIS_PASSWORD|String|redispass|||
|REDIS_DB|Integer|0|||
|REDIS_POOL_SIZE|Integer|20|true||
|REDIS_DIAL_TIMEOUT|Duration|5s|true||
|REDIS_READ_WRITE_TIMEOUT|Duration|3s|true||
|REDIS_IDLE_TIMEOUT|Duration|5m|true||
|REDIS_IDLE_CHECK_FREQUENCY|Duration|1m|true||
|REDIS_MAX_CONN_AGE|Duration|30m|true||

