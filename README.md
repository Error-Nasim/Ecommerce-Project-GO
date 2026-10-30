# E-Commerce REST API in Go 🛒🐹

A modular, lightweight RESTful E-Commerce API built with **Go (Golang)** using the standard library `net/http`, **PostgreSQL** with `sqlx`, and clean architectural patterns based on **Domain-Driven Design (DDD)** and **Ports & Adapters (Hexagonal Architecture)**.

This project is created as part of my journey learning backend development in Go, following the **[Go with Habib - Golang Backend Development (Beginner Friendly)](https://youtube.com/playlist?list=PLpCqPSEm2Xe8sEY2haMDUVgwbkIs5NCJI)** YouTube channel tutorials.

---

## 📌 Project Overview

This API provides core e-commerce backend functionalities including product management, user registration, and authentication using custom JSON Web Tokens (JWT). 

The codebase follows a clean, feature-driven modular architecture with **Domain-Driven Design (DDD)**, service orchestration via **Ports & Adapters**, the **Repository Pattern (Go interfaces)**, PostgreSQL persistence via `sqlx`, automated database migrations via `sql-migrate`, dependency injection, standard library routing (`net/http.ServeMux`), custom middleware chaining, and dedicated request/response DTOs (including pagination)—all without relying on heavy third-party web frameworks.

---

## ✨ Features & Architecture

- **Domain-Driven Design (DDD) & Ports & Adapters Architecture**:
  - **Domain Entities (`domain/`)**: Pure core business models (`domain.Product`, `domain.User`) decoupled from HTTP transport and database concerns.
  - **Domain Services & Ports (`product/`, `user/`)**:
    - Centralizes business logic in dedicated service implementations (`product.Service`, `user.Service`).
    - Declares repository interfaces (`ProductRepo`, `UserRepo`) as secondary/driven ports.
  - **Inverted HTTP Service Ports (`rest/handlers/*/port.go`)**:
    - Handlers define the contract (`Service` interface) they require from the business layer, achieving true decoupling and inversion of control.
  - **Infrastructure Adapters (`repo/`)**:
    - PostgreSQL repository implementations satisfying domain repository ports.
- **Full Dependency Injection Chain (`cmd/serve.go`)**:
  - `Database Connection` -> Injected into `Repositories` (`repo.NewProductRepo`, `repo.NewUserRepo`).
  - `Repositories` -> Injected into `Domain Services` (`product.NewService`, `user.NewService`).
  - `Domain Services` -> Injected into `HTTP Handlers` (`prdctHandler.NewHandler`, `usrHandler.NewHandler`).
  - `HTTP Handlers` -> Injected into `Server` (`rest.NewServer`).
- **Persistent Database (PostgreSQL & `sqlx`)**:
  - Integrated with PostgreSQL using `github.com/jmoiron/sqlx` and `github.com/lib/pq`.
  - Full CRUD operations for Products (`Create`, `Get`, `List`, `Update`, `Delete`) and Users (`Create`, `Find`).
  - Safe parameterized SQL queries (`$1, $2, ...`) and `RETURNING id` statements.
  - Dynamic connection builder in `infra/db/connection.go` assembling connection strings from structured configuration.
- **Automated Database Migrations**:
  - Powered by `github.com/rubenv/sql-migrate` in `infra/db/migrate.go`.
  - Versioned migration files in `migrations/` with `-- +migrate Up` and `-- +migrate Down` definitions.
  - Automatically runs pending migrations on application startup (`cmd/serve.go`).
- **Pagination & Query Parameters**:
  - Products listing supports `page` and `limit` query parameters with safe fallback defaults (`page=1`, `limit=10`).
  - Repository calculates dynamic SQL `LIMIT` and `OFFSET` queries alongside `SELECT COUNT(*)` for total items.
  - Standardized JSON pagination envelope formatted with `util.SendPage`: includes current page, limit, total items, and total pages.
- **Centralized Configuration & Secrets Management**:
  - Strongly typed configuration structs (`config.Config`, `config.DBConfig`) parsing server and database options.
  - Secure `.env` handling: `.env` is untracked in `.gitignore`, and a clean [`.env.example`](.env.example) template is provided for setup.
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
│   └── serve.go                  # Application bootstrapper, DB init, auto-migration & full DI wiring
├── config/
│   └── config.go                 # Environment & DB config loader (Config, DBConfig)
├── db_queries/                   # Standalone reference SQL DDL schemas
│   ├── 001-user.sql              # Users table schema
│   └── 002-create-table-product.sql # Products table schema
├── domain/                       # Core Domain Entities (Layer 1)
│   ├── product.go                # Pure Product domain model
│   └── user.go                   # Pure User domain model
├── infra/                        # Infrastructure & Database Setup
│   └── db/
│       ├── connection.go         # PostgreSQL connection builder using sqlx & lib/pq
│       └── migrate.go            # Migration runner using sql-migrate
├── migrations/                   # Versioned SQL migration files
│   ├── 000001-create-users.up.sql
│   ├── 000001-create-users.down.sql
│   ├── 000002-create-products.up.sql
│   └── 000002-create-products.down.sql
├── product/                      # Product Domain Service Layer (Layer 2)
│   ├── port.go                   # Product Service & ProductRepo interfaces (Ports)
│   └── service.go                # Product business logic service implementation
├── repo/                         # Repository Adapters (Layer 3)
│   ├── product.go                # Product repository implementing product.ProductRepo with PostgreSQL & pagination
│   └── user.go                   # User repository implementing user.UserRepo with PostgreSQL
├── rest/                         # HTTP Transport & Delivery (Layer 3)
│   ├── handlers/
│   │   ├── product/              # Product HTTP handler module
│   │   │   ├── handler.go        # Handler struct receiving middlewares & product Service port
│   │   │   ├── port.go           # Service contract required by Product handler (List with pagination, Count)
│   │   │   ├── routes.go         # Product routes registration
│   │   │   ├── create_product.go # Create new product (Protected, uses ReqCreateProduct DTO)
│   │   │   ├── get_products.go   # List products with page & limit query parameters
│   │   │   ├── get_product.go    # Fetch product by ID
│   │   │   ├── update_product.go # Update existing product (Protected, uses ReqUpdateProduct DTO)
│   │   │   └── delete_product.go # Delete product by ID (Protected)
│   │   └── user/                 # User HTTP handler module
│   │       ├── handler.go        # Handler struct receiving config & user Service port
│   │       ├── port.go           # Service contract required by User handler
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
├── user/                         # User Domain Service Layer (Layer 2)
│   ├── port.go                   # User Service & UserRepo interfaces (Ports)
│   └── service.go                # User business logic service implementation
├── util/
│   ├── create_jwt.go             # HMAC-SHA256 JWT creator & Base64URL encoder
│   ├── send_data.go              # JSON response helpers (SendData, SendError)
│   └── send_page.go              # Standardized JSON pagination response helper (SendPage, PaginatedData)
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
4. Wire together dependencies (Repositories -> Domain Services -> HTTP Handlers -> Server).
5. Start the HTTP server on `:4000`.

```text
Sucessfully migrated database
Server running on port :4000
```

---

## 📡 API Endpoints

### 🔐 Public Endpoints

| Method | Endpoint | Description | Auth Required |
|---|---|---|---|
| `GET` | `/products?page=1&limit=10` | List products with pagination | No |
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

### 4. Get Paginated Products (`GET /products?page=1&limit=10`)

**Request:**
```http
GET /products?page=1&limit=10 HTTP/1.1
```

**Response (200 OK):**
```json
{
  "data": [
    {
      "id": 1,
      "title": "Watermelon",
      "description": "Fresh and sweet organic watermelon",
      "price": 120,
      "imageUrl": "https://example.com/watermelon.jpg"
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "totalItems": 1,
    "totalPages": 0
  }
}
```

---

## 🗺️ Learning Roadmap & Course Progress

Follows the **[Go with Habib - Golang Backend Development (Beginner Friendly)](https://youtube.com/playlist?list=PLpCqPSEm2Xe8sEY2haMDUVgwbkIs5NCJI)** playlist:

### 🟢 Completed Milestones (Videos 047 – 072)

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
  - [x] Defined repository interfaces
  - [x] Constructor dependency injection
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
- [x] **Domain-Driven Design (DDD) & Ports & Adapters Refactoring** *(Videos 067–069)*
  - [x] Centralized pure business entities in `domain/` (`Product`, `User`)
  - [x] Domain Service layer in `product/` and `user/` handling business logic
  - [x] Inverted service ports (`rest/handlers/*/port.go`) decoupling HTTP handlers from services
  - [x] Clean multi-tiered Dependency Injection wiring in `cmd/serve.go`
- [x] **Pagination & Query Parameters** *(Videos 071–072)*
  - [x] Query parameter extraction for `page` and `limit` in `GetProducts`
  - [x] SQL pagination queries (`LIMIT $1 OFFSET $2`) and `SELECT COUNT(*)` in repository
  - [x] Standardized `util.SendPage` JSON helper with `PaginatedData` and `Pagination` metadata structs

---

### 🟡 Upcoming Milestones (Videos 070, 073 – 079)

- [ ] **Phase 6: Multilingual Support** *(Video 070)*
  - [ ] **070: Internationalization**: Multilingual support (i18n)
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
