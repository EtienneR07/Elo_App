<script lang="ts">
  import { onMount } from 'svelte';
  import { navigate } from '$app/navigation';
  import { auth } from '$lib/stores/auth';

  onMount(async () => {
    const hashParams = new URLSearchParams(window.location.hash.split('?')[1]);
    const searchParams = new URLSearchParams(window.location.search);
    const token = hashParams.get('token') || searchParams.get('token');

    if (token) {
      try {
        await auth.handleOAuthCallback(token);
        navigate('/');
      } catch (error) {
        console.error('OAuth error:', error);
        navigate('/login');
      }
    } else {
      navigate('/login');
    }
  });
</script>
