<script lang="ts">
  import { onMount } from 'svelte';
  import { initRouter, getCurrentPath, matchRoute, type Route } from './lib/router.svelte';
  import Home from './pages/Home.svelte';
  import Login from './pages/Login.svelte';
  import Register from './pages/Register.svelte';
  import AuthCallback from './pages/AuthCallback.svelte';

  const routes: Route[] = [
    { path: '/', component: Home },
    { path: '/login', component: Login },
    { path: '/register', component: Register },
    { path: '/auth/callback', component: AuthCallback }
  ];

  onMount(() => {
    initRouter();
  });

  let Component = $derived(matchRoute(routes));
</script>

{#if Component}
  <Component />
{/if}
