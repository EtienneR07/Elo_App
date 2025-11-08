const http = require('http');

const PORT = 3001;
const HOST = 'localhost';

const server = http.createServer((req, res) => {
  // Set CORS headers to allow frontend requests
  res.setHeader('Access-Control-Allow-Origin', '*');
  res.setHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE');
  res.setHeader('Access-Control-Allow-Headers', 'Content-Type');

  // Handle different routes
  if (req.url === '/' && req.method === 'GET') {
    res.statusCode = 200;
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify({
      message: 'Hello from Elo App Backend!',
      timestamp: new Date().toISOString()
    }));
  } else if (req.url === '/api/status' && req.method === 'GET') {
    res.statusCode = 200;
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify({
      status: 'running',
      uptime: process.uptime()
    }));
  } else {
    res.statusCode = 404;
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify({ error: 'Not Found' }));
  }
});

server.listen(PORT, HOST, () => {
  console.log(`Server running at http://${HOST}:${PORT}/`);
  console.log('Available endpoints:');
  console.log(`  - http://${HOST}:${PORT}/`);
  console.log(`  - http://${HOST}:${PORT}/api/status`);
});
