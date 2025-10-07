```markdown
# 🧩 Golang Microservice - Clean Architecture CRUD API

A fully modular implementation of **Clean Architecture** in Go using **Gin Framework**, **PostgreSQL**, and **Zap Logger**.  
This project is designed to be testable, maintainable, and production-ready — serving as a foundation for scalable microservices.

---

## 🚀 Features

- ✅ Clean Architecture (entity → usecase → repository → delivery)
- 🧠 Full CRUD operations (Create, Read, Update, Delete)
- 🔐 API Key middleware (`X-API-Key`)
- 🧾 Zap structured logging (stdout)
- 🐘 PostgreSQL integration
- 🧱 Docker multi-stage build ready
- 🧰 Environment configuration via `.env`
- 🧪 Postman collection included

---

## 🧱 Project Structure

```bash
golang-microservice/
├── cmd/
│   └── server/
│       └── main.go                  # 🏁 Entry point (bootstrap app)
│
├── internal/
│   ├── entity/                      # 🧠 Domain entities
│   │   └── user.go
│   │
│   ├── usecase/                     # 💡 Business logic
│   │   └── user_usecase.go
│   │
│   ├── repository/                  # 💾 Data access implementation
│   │   └── user_repository.go
│   │
│   ├── delivery/                    # 🌐 Interface Adapters (Gin handlers)
│   │   └── http/
│   │       ├── handler/
│   │       │   └── user_handler.go
│   │       └── router.go
│   │
│   ├── infrastructure/              # ⚙️ Frameworks & drivers
│   │   ├── db/
│   │   │   └── postgres.go
│   │   ├── logger/
│   │   │   └── zap_logger.go
│   │   ├── config/
│   │   │   └── config.go
│   │   └── middleware/
│   │       └── auth.go
│   │
│   ├── app/                         # 🧩 Dependency injection (layer wiring)
│   │   └── app.go
│   │
│   └── utils/                       # 🛠️ Common utilities
│       └── response.go
│
├── pkg/                             # 📦 Optional reusable packages
│   └── pagination/
│       └── paginate.go
│
├── test/                            # 🧪 Unit / integration tests
│   └── user_usecase_test.go
│
├── go.mod
├── go.sum
├── Dockerfile
└── README.md


---

## ⚙️ Environment Setup

### 1️⃣ Clone the repository

```bash
git clone https://github.com/ekobudiarto/golang-microservice.git
cd golang-microservice
````

### 2️⃣ Create a `.env` file

```env
APP_PORT=8080
API_KEY=my-secret-key

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=userdb
```

### 3️⃣ Setup PostgreSQL database

```sql
CREATE DATABASE userdb;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE
);
```

---

## 🧩 Run the Application

### Run locally

```bash
go mod tidy
go run ./cmd/server
```

Server will be running at:

```
http://localhost:8080/api/v1
```

### Run with Docker

```bash
docker build -t golang-microservice .
docker run -p 8080:8080 --env-file .env golang-microservice
```

---

## 🧠 API Endpoints

All endpoints require the following header:

```
X-API-Key: my-secret-key
```

| Method   | Endpoint     | Description    |
| -------- | ------------ | -------------- |
| `POST`   | `/users`     | Create user    |
| `GET`    | `/users`     | Get all users  |
| `GET`    | `/users/:id` | Get user by ID |
| `PUT`    | `/users/:id` | Update user    |
| `DELETE` | `/users/:id` | Delete user    |

---

## 🧾 Example Requests (cURL)

### Create User

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "X-API-Key: my-secret-key" \
  -H "Content-Type: application/json" \
  -d '{"name": "Eko Budiarto", "email": "eko@example.com"}'
```

### Get All Users

```bash
curl -X GET http://localhost:8080/api/v1/users \
  -H "X-API-Key: my-secret-key"
```

### Update User

```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "X-API-Key: my-secret-key" \
  -H "Content-Type: application/json" \
  -d '{"name": "Eko Updated", "email": "eko.updated@example.com"}'
```

### Delete User

```bash
curl -X DELETE http://localhost:8080/api/v1/users/1 \
  -H "X-API-Key: my-secret-key"
```

---

## 🧪 Postman Collection

Ready-to-use Postman collection:
📁 [`golang_microservice_collection.json`](./golang_microservice_collection.json)

After importing, set your Postman environment variables:

| Variable   | Value                          |
| ---------- | ------------------------------ |
| `base_url` | `http://localhost:8080/api/v1` |
| `api_key`  | `my-secret-key`                |

---

## 🧰 Logging

This project uses **Zap Logger** for structured JSON logs written to stdout.

Example log output:

```json
{"level":"info","msg":"Server running on port 8080"}
{"level":"info","msg":"POST /api/v1/users"}
```

---

## 🧠 Clean Architecture Overview

```text
[Entity] <-> [Usecase] <-> [Repository] <-> [Infrastructure]
                      ^
                      |
                  [Delivery (HTTP)]
```

* **Entity** → domain models
* **Usecase** → core business rules
* **Repository** → data access layer (PostgreSQL)
* **Delivery** → HTTP handler layer (Gin)
* **Infrastructure** → framework, config, and middleware

---

## 🧱 Build & Deployment

### Build binary

```bash
go build -o server ./cmd/server/main.go
```

### Run binary

```bash
./server
```

### Build Docker image (multi-stage)

```bash
docker build -t golang-microservice .
```

### Run container

```bash
docker run -p 8080:8080 golang-microservice
```

---

## 🧪 Testing

To run unit or integration tests:

```bash
go test ./test/...
```

---

## 🧩 License

MIT License © 2025 [Eko Budiarto](https://github.com/ekobcode)