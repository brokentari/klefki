<script>
	import { onMount } from 'svelte';
	import Navbar from '../Home/Navbar.svelte';
	import PasswordDialog from './PasswordDialog.svelte';

	/** @type {import('$lib/types').User}*/
	export let current_user;
	/** @type boolean*/
	export let loggedIn;

	/** @type {import('$lib/types').Password[]}*/
	let user_passwords;
	/** @type {import('$lib/types').Password}*/
	let selected_password;
	/** @type boolean */
	let hidden = true;

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

			current_user = /** @type {import('$lib/types').User}*/ (responseJson);

			let passwords_response = await fetch(
				'http://localhost:3000/password?' +
					new URLSearchParams({
						userId: storedUser
					})
			);
			let password_json = await passwords_response.json();
			user_passwords = /** @type {import('$lib/types').Password[]}*/ (password_json);

			console.log(user_passwords);
		}
	});

	/**
	 * Handle event when a password is selected. Opens a dialog
	 * @param {import('$lib/types').Password} password
	 */
	const handleOpenPassword = (password) => {
		console.log('clicked on a password');
		selected_password = password;
		hidden = false;
	};
</script>

<Navbar bind:current_user bind:loggedIn />
<div class="pt-20">
	<PasswordDialog bind:selected_password bind:hidden />

	{#if user_passwords}
		{#each user_passwords as pass}
			<!-- svelte-ignore a11y-no-static-element-interactions -->
			<div
				on:click={() => handleOpenPassword(pass)}
				on:keypress={() => handleOpenPassword(pass)}
				class="bg-white rounded-lg shadow-md p-4 cursor-pointer m-4"
				id="card"
			>
				<div class="flex items-center">
					<div class="bg-blue-500 rounded-full p-2 mr-4">
						<svg class="w-6 h-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"
							/>
						</svg>
					</div>
					<div>
						<p class="text-xl font-semibold text-gray-800 hover:text-blue-500">
							{pass.name}
						</p>
						<p class="text-gray-500">{pass.username}</p>
					</div>
				</div>
			</div>
		{/each}
	{/if}
</div>

<button
	class="fixed bottom-4 right-4 bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-full shadow-md"
>
	add</button
>
