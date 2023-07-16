<script>
	import Home from '$lib/Home/Home.svelte';
	import LoginForm from '$lib/Auth/LoginForm.svelte';
	import { onMount } from 'svelte';

	let loggedIn = false;
	let checkedLoginStatus = false;

	/** @type {import('$lib/types').User} */
	let current_user;

	onMount(() => {
		const storedUser = localStorage.getItem('userId');
		if (storedUser) {
			loggedIn = true;
		}
		checkedLoginStatus = true;
	});
</script>

<div class="grid gap- h-screen place-items-center">
	{#if loggedIn}
		<Home bind:current_user bind:loggedIn />
	{:else if checkedLoginStatus}
		<LoginForm bind:loggedIn bind:current_user />
	{/if}
</div>

<style lang="postcss">
	:global(html) {
		background-color: theme('colors.gray.900');
	}
</style>
