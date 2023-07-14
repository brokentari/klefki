<script>
	import { onMount } from 'svelte';
	import Navbar from '../Home/Navbar.svelte';

	/** @type {import('../../routes/+page.svelte').User}*/
	export let current_user;
	/** @type boolean*/
	export let loggedIn;

	onMount(async () => {
		let storedUser = localStorage.getItem('userId');

		if (storedUser) {
			let response = await fetch(
				'http://localhost:3000/profile?' +
					new URLSearchParams({
						userId: storedUser
					})
			);
			let responseJson = await response.json();

			current_user = /** @type {import('../../routes/+page.svelte').User}*/ (responseJson);
		}
	});

	const handleLogout = () => {
		localStorage.removeItem('userId');
	};
</script>

<Navbar bind:current_user bind:loggedIn />
