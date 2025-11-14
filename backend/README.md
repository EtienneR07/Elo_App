# Elo App Backend

Go backend API using Gin framework with JWT authentication.

## Prerequisites

- Go 1.21 or higher
- Air (for hot reload): `go install github.com/air-verse/air@latest`

## Installation

```bash
cd backend
go mod tidy
```

## Running the Server

### Option 1: With Hot Reload (Recommended for Development)

```bash
cd backend
air
```

This will:
- Auto-reload when you save Go files
- Only create one firewall rule (uses tmp/main.exe)
- Fast rebuild on changes

### Option 2: Direct Run

```bash
cd backend
go run main.go
```

### Option 3: Build and Run

```bash
cd backend
go build -o elo-backend.exe
.\elo-backend.exe
```

## WebStorm/GoLand

Two run configurations are available:

1. **Backend with Air** - Hot reload development
2. **Backend Debug** - Full debugging with breakpoints

Access them from the run configurations dropdown in the top toolbar.

## API Endpoints

### Authentication (Public)

- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login
- `POST /api/auth/logout` - Logout

### Protected Routes (Requires JWT)

- `GET /api/auth/session` - Get current user session
- `GET /api/leagues` - Get all leagues
- `POST /api/leagues` - Create league
- `GET /api/leagues/:id` - Get specific league
- `PUT /api/leagues/:id` - Update league
- `DELETE /api/leagues/:id` - Delete league

## Default Test Account

```
Email: test@example.com
Password: password
```

## JWT Authentication

Include the JWT token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

## Tech Stack

- Gin - HTTP web framework
- golang-jwt/jwt - JWT authentication
- bcrypt - Password hashing
- Air - Hot reload (development)
- CORS enabled for frontend (localhost:5173, localhost:5174)
