# E-Commerce REST API in Go 🛒🐹

A modular, lightweight RESTful E-Commerce API built with **Go (Golang)** using the standard library `net/http` and clean architectural patterns.

This project is created as part of my journey learning backend development in Go, following the **[Go with Habib](https://www.youtube.com/@gowithhabib)** YouTube channel tutorials.

---

## 📌 Project Overview

This API provides core e-commerce backend functionalities including product management, user registration, and authentication using custom JSON Web Tokens (JWT). The codebase is designed to avoid unnecessary heavy frameworks by leveraging Go's modern standard library routing (`net/http.ServeMux`), custom middleware chaining, and in-memory data structures.

---

## ✨ Features

- **Standard Library Routing**: Uses Go 1.22+ enhanced `net/http.ServeMux` with pattern matching (`METHOD /path/{id}`).
- **Custom Middleware Architecture**: A reusable `Manager` to handle global and route-specific middleware pipelines.
  - **Logging**: Request method, URI, status code, and duration tracking.
  - **CORS & Preflight**: Configured HTTP headers and `OPTIONS` request handling.
  - **JWT Authentication**: Protected routes requiring valid Bearer tokens.
- **Custom JWT Implementation**: Hand-crafted HS256 (HMAC-SHA256) token generation and verification without external JWT libraries.
- **Product Management**: Full CRUD operations for products (in-memory mock database).
- **User Authentication**: User registration and login with JWT access token generation.
- **Environment Configuration**: Centralized environment variable loader using `godotenv`.

---

## 📂 Project Structure

```text
├── cmd/
│   └── serve.go                  # Server command runner
├── config/
│   └── config.go                 # Environment config loader (.env validation)
├── database/
│   ├── product.go                # Product model & in-memory data store with seed items
│   └── user.go                   # User model & in-memory data store
├── rest/
│   ├── handlers/
│   │   ├── create_product.go     # Create new product
│   │   ├── get_products.go       # List all products
│   │   ├── get_product.go        # Fetch product by ID
│   │   ├── update_product.go     # Update existing product
│   │   ├── delete_product.go     # Delete product by ID
│   │   ├── create_user.go        # Register new user
│   │   └── login.go              # User login & JWT issuance
│   ├── middlewares/
│   │   ├── manager.go            # Middleware manager & chaining helper
│   │   ├── logger.go             # HTTP request logger
│   │   ├── cors.go               # CORS headers middleware
│   │   ├── preflight.go          # OPTIONS preflight request handler
│   │   └── authenticate_jwt.go   # JWT validation middleware
│   ├── routes.go                 # Route definitions & middleware attachment
│   └── server.go                 # Server initialization & HTTP listener
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

The server should start on the configured port:
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

**Response (200 OK):**
```json
[
  {
    "id": 1,
    "title": "Orange",
    "description": "Very Sweet Orange",
    "price": 100,
    "imageUrl": "..."
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
- [ ] Connect a persistent database (PostgreSQL / MySQL)
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
