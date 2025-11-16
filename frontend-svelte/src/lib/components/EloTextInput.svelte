<script lang="ts">
    export let disabled = false;
    export let value = '';
    export let label = '';
    export let type: 'email' | 'password' | 'text' = 'text';
    export let required = false;

    $: emailInputClasses = [
        'peer block w-full px-4 pt-6 pb-2 border border-elo-primary-500 rounded-xl focus:outline-none focus:ring-2 focus:ring-elo-primary-500 focus:border-elo-primary-500 placeholder-transparent',
        disabled && 'opacity-50 cursor-not-allowed'
    ].filter(Boolean).join(' ');

    $: labelClasses = [
        'absolute font-medium left-4 top-2 text-xs text-gray-500 transition-all duration-200',
        'peer-placeholder-shown:top-1/2 peer-placeholder-shown:-translate-y-1/2 peer-placeholder-shown:text-base peer-placeholder-shown:text-elo-primary-500',
        'peer-focus:top-2 peer-focus:translate-y-0 peer-focus:text-xs peer-focus:text-elo-primary-500',
        disabled && 'opacity-50'
    ].filter(Boolean).join(' ');
</script>

<div class="relative">
    <input
        type={type}
        bind:value
        class={emailInputClasses}
        disabled={disabled}
        required={required}
        placeholder={label}
    />
    {#if label}
        <label for={label} class={labelClasses}>
            {label}
        </label>
    {/if}
</div>

