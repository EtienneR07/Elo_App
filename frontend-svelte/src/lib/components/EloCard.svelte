<script lang="ts">
    export let label = '';
</script>

<div class="max-w-md w-full p-8 bg-white shadow-2xl z-10 rounded-xl">
    <div>
        <h2 class="text-center text-3xl font-medium text-elo-primary-500">{label}</h2>
    </div>

    <slot />
</div>

