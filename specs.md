# Todo RESTful API Specification

## Project Overview

Membuat RESTful API sederhana untuk aplikasi Todo List menggunakan:

- Golang
- HttpRouter (github.com/julienschmidt/httprouter)
- MySQL
- database/sql
- JSON sebagai format request dan response

Project ini dibuat sebagai latihan backend yang mengikuti praktik yang baik namun tetap mudah dipahami oleh junior programmer maupun AI coding agent.

---

# Tujuan Project

Membuat API yang dapat melakukan:

- Membuat Todo
- Melihat seluruh Todo
- Melihat detail Todo
- Mengubah Todo
- Menghapus Todo
- Menandai Todo selesai
- Menandai Todo belum selesai

Tidak perlu autentikasi.

---

# Tech Stack

| Technology | Version |
|------------|----------|
| Go | 1.24+ |
| HttpRouter | Latest |
| MySQL | 8+ |
| database/sql | Standard Library |
| go-sql-driver/mysql | Latest |

---

# Project Structure

\`\`\`
todolist-api/
├── cmd/
│   └── server/
│       └── main.go
├── config/
│   └── database.go
├── controllers/
│   └── todo_controller.go
├── models/
│   └── todo.go
├── repositories/
│   └── todo_repository.go
├── services/
│   └── todo_service.go
├── routes/
│   └── route.go
├── helpers/
│   ├── response.go
│   └── error.go
├── migrations/
│   └── schema.sql
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
\`\`\`

---

# Architecture

Request → Router → Controller → Service → Repository → MySQL

Tanggung jawab setiap layer:

- Controller: parsing request & response.
- Service: business logic dan validasi.
- Repository: query database.
- Model: representasi data.

---

# Database

Database:

\`\`\`
todo_db
\`\`\`

Table:

\`\`\`
todos
\`\`\`

Schema:

\`\`\`sql
CREATE TABLE todos (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        ON UPDATE CURRENT_TIMESTAMP
);
\`\`\`

---

# Model

\`\`\`go
type Todo struct {
    ID          int64
    Title       string
    Description string
    Completed   bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
\`\`\`

---

# Environment

\`\`\`
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=todo_db
SERVER_PORT=8080
\`\`\`

---

# API Base URL

\`\`\`
/api/v1
\`\`\`

Semua response menggunakan:

- Content-Type: application/json

---

# Response Format

Success

\`\`\`json
{
  "success": true,
  "message": "Success",
  "data": {}
}
\`\`\`

Error

\`\`\`json
{
  "success": false,
  "message": "Validation Error",
  "errors": [
    "title is required"
  ]
}
\`\`\`

---

# Endpoint

## GET /todos

Mengambil seluruh todo.

## GET /todos/:id

Mengambil detail todo.

404 jika tidak ditemukan.

## POST /todos

Request:

\`\`\`json
{
  "title": "Belajar Golang",
  "description": "Belajar REST API"
}
\`\`\`

Validasi:

- title wajib
- minimal 3 karakter
- maksimal 255 karakter

Response:

201 Created

---

## PUT /todos/:id

Mengubah seluruh data.

---

## PATCH /todos/:id/complete

completed = true

---

## PATCH /todos/:id/uncomplete

completed = false

---

## DELETE /todos/:id

204 No Content

---

# HTTP Status

- 200 OK
- 201 Created
- 204 No Content
- 400 Bad Request
- 404 Not Found
- 500 Internal Server Error

---

# Validation

Title

- Required
- Min 3 karakter
- Max 255 karakter

Description

- Optional
- Max 1000 karakter

ID

- Integer positif

---

# Repository Contract

\`\`\`go
FindAll()
FindByID(id)
Create(todo)
Update(todo)
Delete(id)
MarkComplete(id)
MarkIncomplete(id)
\`\`\`

---

# Service Responsibility

- Validasi
- Business Logic
- Memanggil Repository
- Mengembalikan error bila data tidak ditemukan

---

# Controller Responsibility

- Parsing request
- Decode JSON
- Memanggil Service
- Encode response

Tidak boleh ada SQL.

---

# Router

- GET /api/v1/todos
- GET /api/v1/todos/:id
- POST /api/v1/todos
- PUT /api/v1/todos/:id
- PATCH /api/v1/todos/:id/complete
- PATCH /api/v1/todos/:id/uncomplete
- DELETE /api/v1/todos/:id

---

# Coding Standard

- Gunakan CamelCase.
- Pisahkan layer sesuai arsitektur.
- Jangan gunakan panic selain startup.
- Semua error wajib ditangani.

---

# Logging

Minimal:

- Server Started
- Database Connected
- Request Error
- Database Error

---

# Acceptance Criteria

Project dinyatakan selesai apabila:

- go run ./cmd/server berhasil dijalankan
- go build ./... tanpa error
- Berhasil terkoneksi ke MySQL
- Semua endpoint berjalan
- Semua response mengikuti format standar
- Validasi berjalan
- Error handling berjalan
- Struktur folder sesuai spesifikasi
- Controller → Service → Repository diterapkan
- Dapat diuji menggunakan Postman atau curl

---

# Bonus

- Pagination
- Search
- Sorting
- Dockerfile
- docker-compose.yml
- Makefile
- Unit Test
- Swagger/OpenAPI
- Graceful Shutdown
- Middleware Logging
- Middleware CORS