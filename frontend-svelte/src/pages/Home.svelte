<script lang="ts">
  import { auth, isAuthenticated } from '$lib/stores/auth';
  import { navigate } from '$app/navigation';
  import EloNav from '$lib/components/EloNav.svelte';

  $: if (!$isAuthenticated) navigate('/login');
</script>

{#if $isAuthenticated}
  <div class="min-h-screen bg-gradient-to-br from-elo-primary-100 to-elo-accent-100">
    <EloNav appName="Elotify">
      <div slot="nav-links" class="flex space-x-4">
        <a href="/" class="text-elo-primary-700 hover:text-elo-primary-900 font-medium transition-colors">Home</a>
        <a href="/leagues" class="text-elo-primary-700 hover:text-elo-primary-900 font-medium transition-colors">Leagues</a>
      </div>
    </EloNav>

    <main class="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
      <div class="bg-white rounded-lg shadow-lg p-6 border border-elo-primary-100">
        <h2 class="text-2xl font-bold text-elo-primary-700 mb-4">Welcome to Elotify!</h2>
        <div class="space-y-2">
          <p class="text-gray-700"><strong class="text-elo-primary-500">Name:</strong> {$auth.user?.name}</p>
          <p class="text-gray-700"><strong class="text-elo-primary-500">Email:</strong> {$auth.user?.email}</p>
        </div>
      </div>
    </main>
  </div>
{/if}
