<script lang="ts">
  import { auth } from '$lib/stores/auth';
  import { navigate } from '$app/navigation';
  import EloTextInput from "$lib/components/EloTextInput.svelte";
  import EloButton from "$lib/components/EloButton.svelte";

  let name = '';
  let email = '';
  let password = '';
  let confirmPassword = '';
  let error = '';
  let loading = false;

  async function handleRegister(e: Event) {
    e.preventDefault();
    error = '';

    if (password !== confirmPassword) {
      error = 'Passwords do not match';
      return;
    }

    if (password.length < 6) {
      error = 'Password must be at least 6 characters';
      return;
    }

    loading = true;

    try {
      await auth.register(email, password, name);
      navigate('/');
    } catch (err: any) {
      error = err.message || 'Registration failed';
    } finally {
      loading = false;
    }
  }
</script>

<div class="min-h-screen flex items-center justify-center bg-elo-primary-100 p-4">
  <div class="max-w-md w-full space-y-8 p-8 bg-white rounded-lg shadow">
    <div>
      <h2 class="text-center text-3xl font-medium  text-elo-primary-500">Sign up</h2>
    </div>

    <form class="mt-8 space-y-6" on:submit={handleRegister}>
      {#if error}
        <div class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
          {error}
        </div>
      {/if}

      <div class="space-y-4">
          <EloTextInput
            type="text"
            label="Name"
            required={true}
            bind:value={name}
          />

          <EloTextInput
            type="email"
            label="Email"
            required={true}
            bind:value={email}
          />

          <EloTextInput
            type="password"
            label="Password"
            required={true}
            bind:value={password}
          />
      </div>

      <div class="flex items-center justify-center">
        <EloButton
          type="submit"
          disabled={loading}
        >
          {loading ? 'Creating account...' : 'Sign up'}
        </EloButton>
      </div>

      <div class="text-center text-sm">
        <span class="text-gray-600">Already have an account? </span>
        <a href="/login" class="font-medium text-elo-primary-600 hover:text-elo-primary-500 underline">
          Log in
        </a>
      </div>
    </form>
  </div>
</div>
