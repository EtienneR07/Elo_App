# Elo App

A full-stack Next.js application with TypeScript and NextAuth.js authentication.

## Features

- ✅ Next.js 15 with App Router
- ✅ TypeScript
- ✅ NextAuth.js authentication
- ✅ API routes for backend logic
- ✅ Simple login/logout flow

## Project Structure

```
Elo_App/
├── app/
│   ├── page.tsx              # Home page with auth status
│   ├── layout.tsx            # Root layout
│   ├── providers.tsx         # Session provider wrapper
│   ├── login/
│   │   └── page.tsx          # Login page
│   └── api/
│       ├── auth/
│       │   └── [...nextauth]/
│       │       └── route.ts  # NextAuth.js API routes
│       ├── hello/
│       │   └── route.ts      # Example API route
│       └── status/
│           └── route.ts      # Status API route
├── .env.local                # Environment variables (not in git)
├── .env.example              # Environment variables template
├── package.json
├── tsconfig.json
└── next.config.js
```

## Getting Started

1. **Install dependencies:**
   ```bash
   npm install
   ```

2. **Set up environment variables:**

   The `.env.local` file is already created with a development secret. For production, generate a secure secret:
   ```bash
   openssl rand -base64 32
   ```

   Then update `.env.local` with your secret.

3. **Run the development server:**
   ```bash
   npm run dev
   ```

4. **Open your browser:**
   - Home: [http://localhost:3000](http://localhost:3000)
   - Login: [http://localhost:3000/login](http://localhost:3000/login)

## Authentication

### Demo Credentials

For testing purposes, use these credentials:
- **Username:** `demo`
- **Password:** `password`

### How It Works

This app uses **NextAuth.js** for authentication:

- **Login page**: `/login` - Simple credentials form
- **Protected content**: Home page shows user info when logged in
- **Session management**: Uses JWT strategy for sessions
- **API routes**: `app/api/auth/[...nextauth]/route.ts` handles all auth requests

**IMPORTANT:** The current authentication is for demo purposes only. In production, you should:
- Replace the hardcoded credentials with a real database
- Use proper password hashing (bcrypt, etc.)
- Add proper user management
- Consider OAuth providers (Google, GitHub, etc.)

## API Endpoints

- `GET /api/hello` - Example API endpoint
- `GET /api/status` - Server status endpoint
- `POST /api/auth/signin` - Sign in (handled by NextAuth.js)
- `POST /api/auth/signout` - Sign out (handled by NextAuth.js)
- `GET /api/auth/session` - Get current session (handled by NextAuth.js)

## Tech Stack

- **Framework**: Next.js 15
- **Language**: TypeScript
- **Authentication**: NextAuth.js
- **Styling**: Inline styles (can be replaced with Tailwind, CSS modules, etc.)
