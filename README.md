# Elo App

A full-stack Next.js application with TypeScript.

## Project Structure

```
Elo_App/
├── app/
│   ├── page.tsx          # Home page (frontend)
│   ├── layout.tsx        # Root layout
│   └── api/              # API routes (backend)
│       ├── hello/
│       │   └── route.ts  # GET /api/hello
│       └── status/
│           └── route.ts  # GET /api/status
├── package.json
├── tsconfig.json
└── next.config.js
```

## Getting Started

1. **Install dependencies:**
   ```bash
   npm install
   ```

2. **Run the development server:**
   ```bash
   npm run dev
   ```

3. **Open your browser:**
   - Frontend: [http://localhost:3000](http://localhost:3000)
   - API endpoints:
     - [http://localhost:3000/api/hello](http://localhost:3000/api/hello)
     - [http://localhost:3000/api/status](http://localhost:3000/api/status)

## How It Works

Next.js handles **both frontend and backend** in one app:

- **Frontend**: React components in `app/page.tsx`, `app/layout.tsx`, etc.
- **Backend**: API routes in `app/api/**/route.ts`

No need for separate servers - it's all in one!
