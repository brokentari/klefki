<script>
	/**
	 * @typedef {Object} User
	 * @property {string} email
	 * @property {string} id
	 * @property {string} name
	 * @property {string} password
	 * @property {string} salt
	 *
	 */

	import Home from '$lib/Home/Home.svelte';
	import LoginForm from '$lib/Auth/LoginForm.svelte';
	import { onMount } from 'svelte';

	let loggedIn = false;

	/** @type {User} */
	let current_user;

	onMount(() => {
		const storedUser = localStorage.getItem('userId');
		if (storedUser) {
			loggedIn = true;
		}
	});
</script>

<div class="grid gap- h-screen place-items-center">
	{#if loggedIn}
		<Home bind:current_user />
	{:else}
		<LoginForm bind:loggedIn bind:current_user />
	{/if}
</div>

<style lang="postcss">
	:global(html) {
		background-color: theme('colors.gray.900');
	}
</style>
