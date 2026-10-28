# E-Commerce REST API in Go 🛒🐹

A modular, lightweight RESTful E-Commerce API built with **Go (Golang)** using the standard library `net/http`, **PostgreSQL** with `sqlx`, and clean architectural patterns.

This project is created as part of my journey learning backend development in Go, following the **[Go with Habib - Golang Backend Development (Beginner Friendly)](https://youtube.com/playlist?list=PLpCqPSEm2Xe8sEY2haMDUVgwbkIs5NCJI)** YouTube channel tutorials.

---

## 📌 Project Overview

This API provides core e-commerce backend functionalities including product management, user registration, and authentication using custom JSON Web Tokens (JWT). 

The codebase follows a clean, feature-driven modular architecture with the **Repository Pattern (Go interfaces)**, PostgreSQL persistence via `sqlx`, automated database migrations via `sql-migrate`, dependency injection, standard library routing (`net/http.ServeMux`), custom middleware chaining, and dedicated request DTOs—all without relying on heavy third-party web frameworks.

---

## ✨ Features & Architecture

- **Domain-Driven Modularization**: Handlers and route registrations are grouped by domain (`product`, `user`), each encapsulating its own `RegisterRoutes` method.
- **Repository Pattern with Interfaces**: Decouples data access from HTTP handler logic via Go interfaces:
  - `repo.UserRepo`: User registration and credential lookup (`Create`, `Find`) backed by **PostgreSQL** via `sqlx`.
  - `repo.ProductRepo`: Full CRUD operations (`Create`, `Get`, `List`, `Update`, `Delete`) backed by **PostgreSQL** via `sqlx`.
- **Persistent Database (PostgreSQL & `sqlx`)**:
  - Integrated with PostgreSQL using `github.com/jmoiron/sqlx` and `github.com/lib/pq`.
  - Safe parameterized SQL queries (`$1, $2, ...`) and `RETURNING id` statements.
  - Infrastructure connection builder in `infra/db/connection.go` assembling connection strings dynamically from structured configuration.
- **Automated Database Migrations**:
  - Migration management powered by `github.com/rubenv/sql-migrate` in `infra/db/migrate.go`.
  - Versioned migration files in `migrations/` with `-- +migrate Up` and `-- +migrate Down` definitions.
  - Automatically runs pending migrations on application startup (`cmd/serve.go`).
- **Centralized Configuration & Secrets Management**:
  - Strongly typed configuration structs (`config.Config`, `config.DBConfig`) parsing server and database options.
  - Secure `.env` handling: `.env` is untracked in `.gitignore`, and a clean [`.env.example`](.env.example) template is provided for setup.
- **Dependency Injection**:
  - `cmd/serve.go` wires dependencies: DB connection (`*sqlx.DB`), repositories (`ProductRepo`, `UserRepo`), configuration (`*Config`), and middlewares (`*Middlewares`).
  - `user.Handler` receives `*Config` and `UserRepo`.
  - `product.Handler` receives `*Middlewares` and `ProductRepo`.
  - `rest.Server` receives `*Config` and domain handlers.
- **Request Data Transfer Objects (DTOs)**: Incoming JSON payloads are deserialized into explicit request DTO structs (`ReqCreateProduct`, `ReqUpdateProduct`, `ReqCreateUser`, `ReqLogin`) rather than raw models.
- **Standard Library Routing**: Uses Go 1.22+ enhanced `net/http.ServeMux` with pattern matching (`METHOD /path/{id}`).
- **Custom Middleware Architecture**:
  - Reusable `Manager` for global and route-specific middleware pipelines (`With`, `WrapMux`).
  - **Injected Middleware Pipeline**: `Middlewares` struct with injected `*config.Config` for fast, direct access to secret keys.
  - **Request Logger**: Tracks method, URL path, HTTP status, and request duration.
  - **CORS & Preflight**: Handles cross-origin requests and `OPTIONS` preflight headers.
  - **JWT Authentication**: Validates Bearer tokens on protected endpoints using HMAC-SHA256 signature verification.
- **Custom JWT Implementation**: Hand-crafted HS256 signature generation and verification without external JWT libraries.

---

## 📂 Project Structure

```text
├── cmd/
│   └── serve.go                  # Application bootstrapper, DB connection, auto-migration & DI
├── config/
│   └── config.go                 # Environment & DB config loader (Config, DBConfig)
├── db_queries/                   # Standalone reference SQL DDL schemas
│   ├── 001-user.sql              # Users table schema
│   └── 002-create-table-product.sql # Products table schema
├── infra/
│   └── db/
│       ├── connection.go         # PostgreSQL connection builder using sqlx & lib/pq
│       └── migrate.go            # Migration runner using sql-migrate
├── migrations/                   # Versioned SQL migration files
│   ├── 000001-create-users.up.sql
│   ├── 000001-create-users.down.sql
│   ├── 000002-create-products.up.sql
│   └── 000002-create-products.down.sql
├── repo/
│   ├── product.go                # Product model, ProductRepo interface & PostgreSQL implementation
│   └── user.go                   # User model, UserRepo interface & PostgreSQL implementation
├── rest/
│   ├── handlers/
│   │   ├── product/              # Product domain module
│   │   │   ├── handler.go        # Product Handler struct with injected middlewares & ProductRepo
│   │   │   ├── routes.go         # Product routes registration
│   │   │   ├── create_product.go # Create new product (Protected, uses ReqCreateProduct DTO)
│   │   │   ├── get_products.go   # List all products
│   │   │   ├── get_product.go    # Fetch product by ID
│   │   │   ├── update_product.go # Update existing product (Protected, uses ReqUpdateProduct DTO)
│   │   │   └── delete_product.go # Delete product by ID (Protected)
│   │   └── user/                 # User domain module
│   │       ├── handler.go        # User Handler struct with injected config & UserRepo
│   │       ├── routes.go         # User routes registration
│   │       ├── create_user.go    # Register new user (persists to PostgreSQL)
│   │       └── login.go          # User login & JWT issuance
│   ├── middlewares/
│   │   ├── middleware.go         # Middlewares struct holding config dependencies
│   │   ├── manager.go            # Middleware manager & chaining helper
│   │   ├── logger.go             # HTTP request logger
│   │   ├── cors.go               # CORS headers middleware
│   │   ├── preflight.go          # OPTIONS preflight request handler
│   │   └── authenticate_jwt.go   # JWT validation middleware using injected config
│   └── server.go                 # Server struct with injected handlers & HTTP listener
├── util/
│   ├── create_jwt.go             # HMAC-SHA256 JWT creator & Base64URL encoder
│   └── send_data.go              # JSON response helpers (SendData, SendError)
├── .env.example                  # Environment variables template
├── .gitignore                    # Git ignore file (excludes .env)
├── go.mod                        # Go module definition
├── go.sum                        # Go module checksums
├── main.go                       # Application entry point
└── README.md                     # Project documentation
```

---

## 🚀 Getting Started

### Prerequisites

- **Go**: Version 1.22 or newer installed ([Download Go](https://go.dev/dl/))
- **PostgreSQL**: Version 14 or newer installed and running
- **Git**: Installed and configured

### 1. Clone the Repository

```bash
git clone https://github.com/Error-Nasim/go.git
cd go
```

### 2. Configure Environment Variables

Copy the template `.env.example` file to `.env`:

```bash
cp .env.example .env
```

Open `.env` and fill in your PostgreSQL and server configurations:

```env
VERSION=1.0.0
SERVICE_NAME=ECOMMERCE
HTTP_PORT=4000
JWT_SECRET_KEY=your_secret_key_here

DB_HOST=localhost
DB_PORT=5432
DB_NAME=ecommerce
DB_USER=postgres
DB_PASSWORD=your_postgres_password
DB_ENABLE_SSL_MODE=false
```

Make sure the PostgreSQL database exists:
```bash
psql -U postgres -c "CREATE DATABASE ecommerce;"
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run the Application

```bash
go run main.go
```

On startup, the application will:
1. Load environment configurations.
2. Establish a connection to PostgreSQL.
3. Automatically execute pending database migrations (`migrations/`).
4. Start the HTTP server on `:4000`.

```text
Sucessfully migrated database
Server running on port :4000
```

---

## 📡 API Endpoints

### 🔐 Public Endpoints

| Method | Endpoint | Description | Auth Required |
|---|---|---|---|
| `GET` | `/products` | List all available products | No |
| `GET` | `/products/{id}` | Get product details by ID | No |
| `POST` | `/users` | Register a new user | No |
| `POST` | `/users/login` | Authenticate user and receive JWT | No |

### 🔒 Protected Endpoints (Requires `Authorization: Bearer <token>`)

| Method | Endpoint | Description | Auth Required |
|---|---|---|---|
| `POST` | `/products` | Create a new product | Yes |
| `PUT` | `/products/{id}` | Update an existing product | Yes |
| `DELETE` | `/products/{id}` | Delete a product | Yes |

---

## 📝 Example Requests

### 1. Register User (`POST /users`)

**Request:**
```http
POST /users HTTP/1.1
Content-Type: application/json

{
  "first_name": "Abu",
  "last_name": "Nasim",
  "email": "nasim@example.com",
  "password": "secretpassword",
  "is_shop_owner": true
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "first_name": "Abu",
  "last_name": "Nasim",
  "email": "nasim@example.com",
  "password": "secretpassword",
  "is_shop_owner": true
}
```

---

### 2. User Login (`POST /users/login`)

**Request:**
```http
POST /users/login HTTP/1.1
Content-Type: application/json

{
  "email": "nasim@example.com",
  "password": "secretpassword"
}
```

**Response (201 Created):**
```json
"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsImZpcnN0X25hbWUiOiJBYnUiLCJsYXN0X25hbWUiOiJOYXNpbSIsImVtYWlsIjoibmFzaW1AZXhhbXBsZS5jb20iLCJpc19zaG9wX293bmVyIjp0cnVlfQ.signature..."
```

---

### 3. Create Product (`POST /products`)

**Request:**
```http
POST /products HTTP/1.1
Authorization: Bearer <YOUR_JWT_TOKEN>
Content-Type: application/json

{
  "title": "Watermelon",
  "description": "Fresh and sweet organic watermelon",
  "price": 120,
  "imageUrl": "https://example.com/watermelon.jpg"
}
```

**Response (201 Created):**
```json
{
  "id": 1,
  "title": "Watermelon",
  "description": "Fresh and sweet organic watermelon",
  "price": 120,
  "imageUrl": "https://example.com/watermelon.jpg"
}
```

---

### 4. Get All Products (`GET /products`)

**Request:**
```http
GET /products HTTP/1.1
```

**Response (201 Created):**
```json
[
  {
    "id": 1,
    "title": "Watermelon",
    "description": "Fresh and sweet organic watermelon",
    "price": 120,
    "imageUrl": "https://example.com/watermelon.jpg"
  }
]
```

---

## 🗺️ Learning Roadmap & Course Progress

Follows the **[Go with Habib - Golang Backend Development (Beginner Friendly)](https://youtube.com/playlist?list=PLpCqPSEm2Xe8sEY2haMDUVgwbkIs5NCJI)** playlist:

### 🟢 Completed Milestones (Videos 047 – 066)

- [x] **REST API Foundation & Standard Library Routing** *(Videos 047–051)*
  - [x] Initial models & REST handlers
  - [x] Preflight handling with `OPTIONS` method
  - [x] Path parameters with enhanced `net/http.ServeMux`
- [x] **Middleware Architecture & Authentication** *(Videos 052–056)*
  - [x] Custom Middleware `Manager` (`With`, `WrapMux`)
  - [x] Global Request Logger and CORS middlewares
  - [x] Custom HS256 JWT generation (HMAC-SHA256, Base64URL)
  - [x] Token verification middleware (`AuthenticateJwt`)
- [x] **Decoupling, Interfaces & Repository Pattern** *(Videos 057–058)*
  - [x] Removed tight coupling between handlers and storage
  - [x] Defined `ProductRepo` and `UserRepo` interfaces in `repo/`
  - [x] Constructor dependency injection (`product.NewHandler`, `user.NewHandler`)
  - [x] Explicit Request DTOs (`ReqCreateProduct`, `ReqUpdateProduct`, `ReqCreateUser`, `ReqLogin`)
- [x] **PostgreSQL Integration & CRUD Operations** *(Videos 059–064)*
  - [x] PostgreSQL table creation & schema modeling
  - [x] Database connection setup via `sqlx` and `lib/pq` (`infra/db/connection.go`)
  - [x] `UserRepo` PostgreSQL implementation with parameterized SQL & `NamedQuery`
  - [x] `ProductRepo` PostgreSQL implementation (full `Create`, `Get`, `List`, `Update`, `Delete`)
- [x] **Database Configurations & Migrations** *(Videos 065–066)*
  - [x] Centralized database credentials in `.env` and `config.DBConfig`
  - [x] Untracked `.env`, added `.gitignore` & `.env.example`
  - [x] Automated migrations using `sql-migrate` in `infra/db/migrate.go`
  - [x] Versioned migration files (`migrations/`) executed automatically on server start

---

### 🟡 Upcoming Milestones (Videos 067 – 079)

- [ ] **Phase 5: Domain-Driven Design (DDD)** *(Videos 067–069)*
  - [ ] **067: Introduction to DDD**: Entities, Aggregates, Value Objects, and Services
  - [ ] **068–069: DDD into Code (Part 1 & 2)**: Refactoring code into clean Domain-Driven architecture
- [ ] **Phase 6: Advanced Backend Features** *(Videos 070–072)*
  - [ ] **070: Internationalization**: Multilingual support (i18n)
  - [ ] **071–072: Pagination**: Query parameters (`limit`, `offset`), database pagination queries, and response mapping
- [ ] **Phase 7: Advanced Go Concurrency & Performance** *(Videos 073–079)*
  - [ ] **073: Goroutines & Channels**: Asynchronous processing in backend services
  - [ ] **074–075: `sync.WaitGroup`**: Managing concurrent background tasks and internals
  - [ ] **076–077: Race Conditions & Mutexes**: Thread-safe shared state with `sync.Mutex` and `sync.RWMutex`
  - [ ] **078–079: Deadlocks & Advanced Channels**: Buffered channels, select statements, and deadlock prevention

---

## 📚 Acknowledgements

- Tutorials and curriculum following the **[Go with Habib - Golang Backend Development (Beginner Friendly)](https://youtube.com/playlist?list=PLpCqPSEm2Xe8sEY2haMDUVgwbkIs5NCJI)** YouTube playlist.
- Official [Go Documentation](https://go.dev/doc/).

---

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).
