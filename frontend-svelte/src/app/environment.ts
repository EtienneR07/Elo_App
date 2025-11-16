// Simple polyfill for SvelteKit's $app/environment
export const browser = typeof window !== 'undefined';
