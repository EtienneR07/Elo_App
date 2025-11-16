<script lang="ts">
  export let variant: 'primary' | 'secondary' | 'accent' | 'danger' | 'warning' = 'primary';
  export let disabled = false;
  export let type: 'button' | 'submit' | 'reset' = 'button';


  $: buttonClasses = [
    'font-medium rounded-3xl transition-all px-12 py-3 cursor-pointer flex items-center justify-center gap-3',
    variant === 'primary' && 'bg-elo-primary-500 hover:bg-elo-primary-700 text-white',
    variant === 'secondary' && 'bg-gray-100 hover:bg-gray-300 text-gray-900',
    variant === 'accent' && 'bg-elo-accent-500 hover:bg-elo-accent-700 text-white',
    variant === 'danger' && 'bg-elo-danger-500 hover:bg-elo-danger-700 text-white',
    variant === 'warning' && 'bg-elo-warning-500 hover:bg-elo-warning-700 text-white',
    disabled && 'opacity-50 cursor-not-allowed'
  ].filter(Boolean).join(' ');
</script>

<button
  class={buttonClasses}
  {disabled}
  {type}
  on:click
>
  <slot />
</button>

