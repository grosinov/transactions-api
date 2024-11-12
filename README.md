# Transaction API

App for migrating and retrieving transactions using Go, GORM, and PostgreSQL.

## Prerequisites

- Go 1.23
- PostgreSQL
- Docker

## Commands

> Run every command in project base folder

To run tests:
```shell
go test ./...
```

To run locally with local PostgreSQL db:
```shell
docker compose up -d
```
## API Documentation

### Migrate

#### Request

- **Method**: `POST`
- **URL**: `/migrate`
- **Body**: The request body must contain the CSV file as `multipart/form-data`.

#### Response

```json
[
  {
    "id": 0,
    "user_id": 0,
    "amount": 0.0,
    "datetime": "YYYY-MM-DDThh:mm:ssZ"
  }
]
```

### Get Balance

#### Request

- **Method**: GET
- **URL**: /users/{user_id}/balance
- **Path Parameters**:
  - user_id (required, integer): The unique identifier of the user whose balance is to be fetched.
- **Query Parameters** (optional):
  - from (string): The start date for filtering transactions. Format: yyyy-MM-ddThh:mm:ssZ.
  - to (string): The end date for filtering transactions. Format: yyyy-MM-ddThh:mm:ssZ.

## Relevant information

- Mocks where generated with [gomock](https://github.com/golang/mock).

## Challenge assumptions

- I assumed that transaction ids are unique, since the exercise example showed repeated ids.