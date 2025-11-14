# Elo App

A rating and league management application using the ELO rating system.

## Project Structure

```
.
├── frontend/          # React + Vite frontend
│   ├── src/
│   │   ├── components/  # React components
│   │   ├── contexts/    # Auth context
│   │   ├── pages/       # Page components
│   │   └── utils/       # API utilities
│   └── package.json
├── backend/           # Go + Gin backend API
│   ├── handlers/      # HTTP handlers
│   ├── middleware/    # Auth middleware
│   ├── models/        # Data models
│   ├── utils/         # JWT utilities
│   └── main.go
└── README.md
```

## Tech Stack

### Frontend
- React 19
- Vite 7
- React Router DOM 7
- Tailwind CSS v4
- TypeScript

### Backend
- Go 1.21+
- Gin web framework
- golang-jwt/jwt for authentication
- bcrypt for password hashing
- CORS enabled

## Prerequisites

- Go 1.21 or higher
- Node.js 18+ and npm
- Air (for hot reload): `go install github.com/air-verse/air@latest`

## VSCode Setup (Recommended)

### 1. Install Extensions

Open VSCode and install recommended extensions (VSCode will prompt you):
- Go (golang.go)
- Prettier
- ESLint
- Tailwind CSS IntelliSense

### 2. Install Air for Hot Reload

```bash
go install github.com/air-verse/air@latest
```

### 3. Run the App

#### Option A: Run Full Stack (Backend + Frontend)
Press `Ctrl+Shift+P` → type "Tasks: Run Task" → select "Run Full Stack"

#### Option B: Run Individually

**Backend with Hot Reload:**
- Press `Ctrl+Shift+P` → "Tasks: Run Task" → "Run Backend with Air"

**Frontend:**
- Press `Ctrl+Shift+P` → "Tasks: Run Task" → "Run Frontend"

#### Option C: Debug Backend
- Press `F5` or click "Run and Debug" in sidebar
- Select "Backend Debug"
- Set breakpoints and debug!

## Manual Setup

### 1. Start the Backend

```bash
cd backend
air              # With hot reload (recommended)
# OR
go run main.go   # Without hot reload
```

Backend will run on `http://localhost:8080`

### 2. Start the Frontend

```bash
cd frontend
npm install      # First time only
npm run dev
```

Frontend will run on `http://localhost:5173`

## API Endpoints

### Authentication (Public)
- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login
- `POST /api/auth/logout` - Logout

### Protected Routes (Requires JWT)
- `GET /api/auth/session` - Get current user
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

## Features

✅ JWT-based authentication
✅ User registration and login
✅ League CRUD operations
✅ Protected routes
✅ Hot reload for backend (Air)
✅ Hot reload for frontend (Vite)
✅ Responsive UI with Tailwind CSS
✅ Full debugging support in VSCode
✅ In-memory data storage (can be replaced with database)

## Development Notes

- Frontend uses localStorage for JWT tokens
- Backend uses CORS to allow requests from localhost:5173/5174
- Password hashing with bcrypt
- JWT tokens valid for 24 hours
- Air only creates one firewall rule (uses tmp/main.exe)

## Next Steps

1. Add database integration (PostgreSQL, MySQL, MongoDB)
2. Add player and match management
3. Implement ELO calculation algorithm
4. Add team management
5. Add statistics and leaderboards
