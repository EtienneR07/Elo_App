<script lang="ts">
    import {auth} from '$lib/stores/auth';
    import {navigate} from '$app/navigation';
    import EloButton from './EloButton.svelte';

    // Props
    export let appName = 'Elotify';
    export let showAuth = true;

    function handleLogout() {
        auth.logout();
        navigate('/login');
    }
</script>

<nav class="bg-white shadow-md border-b border-gray-200">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
            <div class="flex items-center">
                <h1 class="text-xl font-bold text-elo-primary-700">{appName}</h1>
            </div>

            <div class="flex items-center">
                <slot name="nav-links"/>
            </div>

            {#if showAuth}
                <div class="flex items-center space-x-4">
                    {#if $auth.user}
                        <span class="text-elo-primary-500 font-medium">Hello, {$auth.user.name}</span>
                        <EloButton variant="danger" on:click={handleLogout}>
                            Logout
                        </EloButton>
                    {:else}
                        <EloButton variant="primary" on:click={() => navigate('/login')}>
                            Sign In
                        </EloButton>
                    {/if}
                </div>
            {/if}
        </div>
    </div>
</nav>
