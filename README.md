# Golang Todo List API

This is my solution to https://roadmap.sh/projects/todo-list-api

## Features

- JWT Autentication: Users can register and login securely
- Todo Management: Create, Read, Update, and Delete from user todos
- Secure user management: user login info is encrypted and stored
- Postgres database: all data is stored in a sql database
- Pagination and filtering: The response is paginated and can be configured using query params

## Running application

### 1. Clone the repo

```bash
git clone https://github.com/umdalecs/todo-list-api
cd todo-list-api
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Environment variables

```bash
cp .env.example .env
```

### 4. Build and run the application

```bash
go build -o out/ ./cmd/api
./out/api
```

## Usage

### Login

```bash
curl --location 'localhost:8080/api/v1/register' \
--header 'Content-Type: application/json' \
--data-raw '{
  "name" :"John Doe",
  "email": "johntest@gmail.com",
  "password": "johndoe123"
}'
```

### Register

It returns an access token

```bash
curl --location 'localhost:8080/api/v1/login' \
--header 'Content-Type: application/json' \
--data '{
  "email": "johntest@gmail.com",
  "password": "johndoe123"
}'
```

### Create Todo

It returns an access token

```bash
curl --location 'localhost:8080/api/v1/todos/' \
--header 'Authorization: Bearer {{token}}' \
--header 'Content-Type: application/json' \
--data '{
    "title": "my super todo item",
    "description": "this is a todo description"
}'
```

### Update Todo


```bash
curl --location --request PUT 'localhost:8080/api/v1/todos/1' \
--header 'Authorization: Bearer {{token}}' \
--header 'Content-Type: application/json' \
--data '{
    "title": "super todo update",
    "description": "this is an update to my super todo"
}'
```

### Delete Todo


```bash
curl --location --request DELETE 'localhost:8080/api/v1/todos/1' \
--header 'Authorization: Bearer {{token}}' \
```

### Get Todos

Use optional query params for pagination and filtering.

```bash
curl --location 'localhost:8080/api/v1/todos?page=1&limit=5&filter=updated' \
--header 'Authorization: Bearer {{token}}'
```

## Todos

- Implement sorting for the to-do list
- Implement unit tests for the API
- Implement rate limiting and throttling for the API
- Implement refresh token mechanism for the authentication