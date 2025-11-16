import { writable, derived } from 'svelte/store';
import { browser } from '$app/environment';

interface User {
  id: string;
  name: string;
  email: string;
}

function createAuthStore() {
  const { subscribe, set } = writable<{
    user: User | null;
    loading: boolean;
  }>({ user: null, loading: true });

  async function initialize() {
    if (!browser) {
      set({ user: null, loading: false });
      return;
    }

    const token = localStorage.getItem('token');
    if (!token) {
      set({ user: null, loading: false });
      return;
    }

    try {
      const payload = JSON.parse(atob(token.split('.')[1]));

      set({
        user: {
          id: payload.userId.toString(),
          email: payload.email,
          name: payload.email.split('@')[0]
        },
        loading: false
      });
    } catch (error) {
      // Invalid token, clear it
      localStorage.removeItem('token');
      set({ user: null, loading: false });
    }
  }

  initialize();

  async function login(email: string, password: string) {
    if (!browser) return;

    try {
      const response = await fetch('http://localhost:8080/api/auth/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password })
      });

      if (!response.ok) {
        const data = await response.json();
        throw new Error(data.error || 'Login failed');
      }

      const data = await response.json();
      localStorage.setItem('token', data.token);
      set({ user: data.user, loading: false });
    } catch (error) {
      throw error;
    }
  }

  async function register(email: string, password: string, name: string) {
    if (!browser) return;

    try {
      const response = await fetch('http://localhost:8080/api/auth/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password, name })
      });

      if (!response.ok) {
        const data = await response.json();
        throw new Error(data.error || 'Registration failed');
      }

      const data = await response.json();
      localStorage.setItem('token', data.token);
      set({ user: data.user, loading: false });
    } catch (error) {
      throw error;
    }
  }

  async function loginWithGoogle() {
    if (!browser) return;
    window.location.href = 'http://localhost:8080/api/auth/google';
  }

  async function handleOAuthCallback(token: string) {
    if (!browser) return;

    try {
      localStorage.setItem('token', token);

      // Decode JWT to get user info (simple decode, not verifying)
      const payload = JSON.parse(atob(token.split('.')[1]));

      set({
        user: {
          id: payload.userId.toString(),
          email: payload.email,
          name: payload.email.split('@')[0] // Use email prefix as name for now
        },
        loading: false
      });
    } catch (error) {
      localStorage.removeItem('token');
      throw error;
    }
  }

  async function logout() {
    if (!browser) return;
    localStorage.removeItem('token');
    set({ user: null, loading: false });
  }

  return {
    subscribe,
    login,
    register,
    loginWithGoogle,
    handleOAuthCallback,
    logout
  };
}

export const auth = createAuthStore();
export const isAuthenticated = derived(auth, $auth => !!$auth.user);
