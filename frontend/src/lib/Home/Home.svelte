<script>
	import { onMount } from 'svelte';
	import Navbar from '../Home/Navbar.svelte';

	/** @type {import('../../routes/+page.svelte').User}*/
	export let current_user;

	onMount(async () => {
		console.log(current_user);
	});
</script>

<Navbar bind:current_user />

<style>
</style>
