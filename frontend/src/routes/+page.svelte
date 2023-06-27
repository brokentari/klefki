<script>
	import LoginForm from '../lib/LoginForm.svelte';
</script>

<div class="grid h-screen place-items-center">
	<LoginForm />
</div>

<style lang="postcss">
	:global(html) {
		background-color: theme('colors.gray.900');
	}
</style>
