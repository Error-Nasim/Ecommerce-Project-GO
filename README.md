# E-Commerce REST API in Go 🛒🐹

A modular, lightweight RESTful E-Commerce API built with **Go (Golang)** using the standard library `net/http` and clean architectural patterns.

This project is created as part of my journey learning backend development in Go, following the **[Go with Habib](https://www.youtube.com/@gowithhabib)** YouTube channel tutorials.

---

## 📌 Project Overview

This API provides core e-commerce backend functionalities including product management, user registration, and authentication using custom JSON Web Tokens (JWT). 

The codebase follows a clean, feature-driven modular architecture with **Repository Pattern (Go interfaces)**, dependency injection, standard library routing (`net/http.ServeMux`), custom middleware chaining, and dedicated request DTOs—all without relying on heavy third-party web frameworks.

---

## ✨ Features & Architecture

- **Domain-Driven Modularization**: Handlers and route registrations are grouped by domain (`product`, `user`), each encapsulating its own `RegisterRoutes` method.
- **Repository Pattern with Interfaces**: Decouples data access from HTTP handler logic via Go interfaces:
  - `repo.ProductRepo`: CRUD operations (`Create`, `Get`, `List`, `Update`, `Delete`) with an in-memory implementation and seed catalog.
  - `repo.UserRepo`: User storage and credential lookup (`Create`, `Find`).
  - Swapping in-memory storage for persistent SQL databases (PostgreSQL/MySQL) requires only a new repository implementation satisfying the interface.
- **Dependency Injection**:
  - `cmd/serve.go` wires dependencies: repositories (`ProductRepo`, `UserRepo`), configuration (`*Config`), and middlewares (`*Middlewares`).
  - `product.Handler` receives `*Middlewares` and `ProductRepo`.
  - `user.Handler` receives `*Config` and `UserRepo`.
  - `rest.Server` receives `*Config` and domain handlers.
- **Request Data Transfer Objects (DTOs)**: Incoming JSON payloads are deserialized into explicit request DTO structs (`ReqCreateProduct`, `ReqUpdateProduct`, `ReqCreateUser`, `ReqLogin`) rather than raw models.
- **Optimized Configuration**: Centralized environment variable loader via `godotenv` with pointer caching (`*Config`) to prevent repetitive reloading of `.env`.
- **Standard Library Routing**: Uses Go 1.22+ enhanced `net/http.ServeMux` with pattern matching (`METHOD /path/{id}`).
- **Custom Middleware Architecture**:
  - Reusable `Manager` for global and route-specific middleware pipelines (`With`, `WrapMux`).
  - **Injected Middleware Pipeline**: `Middlewares` struct with injected `*config.Config` for fast, direct access to secret keys.
  - **Request Logger**: Tracks method, URL path, HTTP status, and request duration.
  - **CORS & Preflight**: Handles cross-origin requests and `OPTIONS` preflight headers.
  - **JWT Authentication**: Validates Bearer tokens on protected endpoints using HMAC-SHA256 signature verification.
- **Custom JWT Implementation**: Hand-crafted HS256 signature generation and verification without external JWT libraries.
- **Product Management**: Full CRUD operations for products with seeded mock data.
- **User Authentication**: Registration and login endpoints returning signed JWT access tokens.

---

## 📂 Project Structure

```text
├── cmd/
│   └── serve.go                  # Application bootstrapper & dependency injection
├── config/
│   └── config.go                 # Environment config loader (.env validation & pointer return)
├── repo/
│   ├── product.go                # Product model, ProductRepo interface & in-memory implementation
│   └── user.go                   # User model, UserRepo interface & in-memory implementation
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
│   │       ├── create_user.go    # Register new user (uses ReqCreateUser DTO)
│   │       └── login.go          # User login & JWT issuance (uses ReqLogin DTO)
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
├── .env                          # Local environment variables
├── go.mod                        # Go module definition
├── go.sum                        # Go module checksums
├── main.go                       # Application entry point
└── README.md                     # Project documentation
```

---

## 🚀 Getting Started

### Prerequisites

- **Go**: Version 1.22 or newer installed ([Download Go](https://go.dev/dl/))
- **Git**: Installed and configured

### 1. Clone the Repository

```bash
git clone https://github.com/Error-Nasim/go.git
cd go
```

### 2. Configure Environment Variables

Create a `.env` file in the root directory:

```env
VERSION=1.0.0
SERVICE_NAME=ECOMMERCE
HTTP_PORT=4000
JWT_SECRET_KEY=your_secret_key_here
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Run the Application

```bash
go run main.go
```

The server will start on the configured port:
```text
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
  "id": 7,
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
    "title": "Orange",
    "description": "Very Sweet Orange, Sweet like Sadie Sink",
    "price": 100,
    "imageUrl": "https://..."
  }
]
```

---

## 🗺️ Learning Roadmap & Progress

- [x] Basic project setup & entry point (`cmd`, `main.go`)
- [x] In-memory database & models (`User`, `Product`)
- [x] CRUD REST handlers for Products
- [x] Standard library multiplexer (`net/http.ServeMux`) with path parameters
- [x] Custom Middleware Manager (`With`, `WrapMux`)
- [x] Logger, CORS, and Preflight middlewares
- [x] User registration & login handlers
- [x] Custom HS256 JWT generator & verification middleware
- [x] **Modular Architecture Refactoring**:
  - [x] Domain-driven handler packages (`product`, `user`)
  - [x] Encapsulated `RegisterRoutes` per domain
  - [x] Dependency injection via `rest.NewServer` and handler constructors
  - [x] Cached config pointer (`*Config`) to avoid repetitive `.env` reading
  - [x] Middlewares struct with injected configuration
- [x] **Repository Pattern & Layer Decoupling**:
  - [x] Segregated data layer into `repo` package with `ProductRepo` and `UserRepo` Go interfaces
  - [x] In-memory repository implementations (`productRepo`, `userRepo`)
  - [x] Handlers depend on repository interfaces rather than concrete package state
  - [x] Dedicated Request DTOs (`ReqCreateProduct`, `ReqUpdateProduct`, `ReqCreateUser`, `ReqLogin`)
  - [x] Injected repositories into handler constructors
- [ ] Connect a persistent database (PostgreSQL / MySQL)
- [ ] Add Reviews domain & handlers
- [ ] Password hashing with `bcrypt`
- [ ] Token expiration and Refresh Token rotation
- [ ] Role-based access control (RBAC - e.g. Shop Owner vs Customer)
- [ ] Input validation library / custom validator
- [ ] Unit & Integration testing

---

## 📚 Acknowledgements

- Tutorials and curriculum inspired by the **[Go with Habib](https://www.youtube.com/@gowithhabib)** YouTube channel.
- Official [Go Documentation](https://go.dev/doc/).

---

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).
