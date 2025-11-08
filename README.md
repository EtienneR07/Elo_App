# Elo App

A full-stack application with Next.js frontend and Node.js backend.

## Project Structure

```
Elo_App/
├── frontend/          # Next.js React application
│   ├── app/          # Next.js app directory
│   ├── package.json
│   └── ...
└── backend/          # Node.js HTTP server
    ├── server.js
    ├── package.json
    └── ...
```

## Getting Started

### Frontend (Next.js)

1. Navigate to the frontend directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Run the development server:
   ```bash
   npm run dev
   ```

4. Open [http://localhost:3000](http://localhost:3000) in your browser

### Backend (Node.js)

1. Navigate to the backend directory:
   ```bash
   cd backend
   ```

2. Install dependencies (if any are added):
   ```bash
   npm install
   ```

3. Start the server:
   ```bash
   npm start
   ```

4. The backend runs at [http://localhost:3001](http://localhost:3001)

   Available endpoints:
   - `GET /` - Hello message
   - `GET /api/status` - Server status

## Running Both Servers

To run both frontend and backend simultaneously, open two terminal windows:

**Terminal 1 (Frontend):**
```bash
cd frontend
npm run dev
```

**Terminal 2 (Backend):**
```bash
cd backend
npm start
```
