_# REST API

<b>This is school project for testing and education purposes - not ready to commercial use.</b>

This API allows you to connect front and backoffice services and perform all operations using one business logic requirements.

# Run API on host

- go run main.go
- visit: http://localhost:8080

# Setup Docker

- docker-compose build
- docker-compose up -d

# PgAdmin

- visit: http://localhost:5050

# TODO:

- structures for config based on .env file,
- better validators and error handling,
- professional auth with ACL,
- logs aggregation Open Telemetry,
- API docs OPEN API 3.0,
- unit tests for requests (on mocks),
- use DRY instead of copy/paste in methods,_
