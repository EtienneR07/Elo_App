// Simple polyfill for SvelteKit's $app/navigation using custom router
import { navigate as routerNavigate } from '../lib/router.svelte';

export const goto = routerNavigate;
export const push = routerNavigate;
export const navigate = routerNavigate;
export const back = () => window.history.back();
export const replace = (url: string) => routerNavigate(url, true);
