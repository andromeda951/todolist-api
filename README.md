# Todo RESTful API

Todo RESTful API menggunakan Golang, HttpRouter, dan MySQL.

## Tech Stack

- Go 1.24+
- HttpRouter
- MySQL 8+
- database/sql
- go-sql-driver/mysql

## Project Structure

```
todolist-api/
├── cmd/server/main.go
├── config/database.go
├── controllers/todo_controller.go
├── helpers/response.go
├── migrations/schema.sql
├── models/todo.go
├── repositories/todo_repository.go
├── routes/route.go
├── services/todo_service.go
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Setup

### 1. Clone Repository

```bash
git clone <repo-url>
cd todolist-api
```

### 2. Setup Database

```bash
mysql -u root -p < migrations/schema.sql
```

### 3. Konfigurasi Environment

Edit file `.env` sesuai environment Anda:

```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=todo_db
SERVER_PORT=8080
```

### 4. Jalankan Aplikasi

```bash
go run ./cmd/server
```

### 5. Build Binary

```bash
go build -o todolist-api ./cmd/server
./todolist-api
```

## API Endpoints

Base URL: `/api/v1`

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /api/v1/todos | Get all todos |
| GET | /api/v1/todos/:id | Get todo by ID |
| POST | /api/v1/todos | Create todo |
| PUT | /api/v1/todos/:id | Update todo |
| PATCH | /api/v1/todos/:id/complete | Mark as completed |
| PATCH | /api/v1/todos/:id/uncomplete | Mark as uncompleted |
| DELETE | /api/v1/todos/:id | Delete todo |

## Contoh Request

### Create Todo

```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Belajar Golang","description":"Belajar REST API"}'
```

### Get All Todos

```bash
curl http://localhost:8080/api/v1/todos
```

### Get Todo by ID

```bash
curl http://localhost:8080/api/v1/todos/1
```

### Update Todo

```bash
curl -X PUT http://localhost:8080/api/v1/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Belajar API","description":"Belajar MySQL","completed":false}'
```

### Mark Complete

```bash
curl -X PATCH http://localhost:8080/api/v1/todos/1/complete
```

### Mark Incomplete

```bash
curl -X PATCH http://localhost:8080/api/v1/todos/1/uncomplete
```

### Delete Todo

```bash
curl -X DELETE http://localhost:8080/api/v1/todos/1
```

## Response Format

Success:

```json
{
  "success": true,
  "message": "Success",
  "data": {}
}
```

Error:

```json
{
  "success": false,
  "message": "Validation Error",
  "errors": ["title is required"]
}