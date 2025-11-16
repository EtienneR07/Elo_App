import { browser } from '$app/environment';

const API_BASE_URL = 'http://localhost:8080';

export async function fetchWithAuth(url: string, options: RequestInit = {}) {
  if (!browser) return new Response();

  const token = localStorage.getItem('token');

  const headers: HeadersInit = {
    'Content-Type': 'application/json',
    ...(options.headers || {})
  };

  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }

  return fetch(`${API_BASE_URL}${url}`, {
    ...options,
    headers
  });
}
