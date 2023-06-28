<script>
	import LoginIcon from './LoginIcon.svelte';
	import RegisterIcon from './RegisterIcon.svelte';

	/**
	 * @enum { String }
	 */
	const State = {
		LOGIN: 'login',
		REGISTER: 'register'
	};

	/** @type {State}*/
	let state = State.LOGIN;
	let email = '';
	let password = '';
	let name = '';

	let loginButtonWidth = 'w-3/4';
	let registerButtonWidth = 'w-1/4';

	/**
	 * Handles a login/register event
	 *
	 * @param e {SubmitEvent} - Form submission
	 * @listens
	 */
	const handleSubmit = async () => {
		if (!email || !password) {
			return;
		}
		let formData = new FormData();

		formData.set('email', email);
		formData.set('password', password);

		if (state === State.REGISTER) {
			formData.set('name', name);
		}

		const response = await fetch('http://localhost:3000/' + state, {
			method: 'POST',
			body: formData
		});

		console.log(response.status);
	};

	/**
	 * Handles state change between registering or logging in
	 *
	 * 	@param new_state {State} - Register/Login button was click
	 *	@listens
	 */
	const handleStateChange = (new_state) => {
		if (new_state != state) {
			name = '';
			email = '';
			password = '';
			switch (new_state) {
				case State.LOGIN:
					state = State.LOGIN;
					loginButtonWidth = 'w-3/4';
					registerButtonWidth = 'w-1/4';
					break;
				case State.REGISTER:
					state = State.REGISTER;
					loginButtonWidth = 'w-1/4';
					registerButtonWidth = 'w-3/4';

					break;
			}
		}
	};
</script>

<div class="bg-sky-950 rounded-2xl pb-12 pt-4 pl-6 pr-6">
	<h3 class="text-center text-white mb-2 font-extrabold">klekfi</h3>

	<form on:submit|preventDefault={handleSubmit} autocomplete="off">
		<div class="mb-4">
			<label for="email" class="block mb-2 text-sm font-medium text-white">Email</label>
			<input
				type="email"
				id="email"
				class="input"
				bind:value={email}
				required
				placeholder="email"
			/>
		</div>

		<div class="mb-6">
			<label for="password" class="block mb-2 text-sm font-medium text-white">Password</label>
			<input
				type="password"
				id="password"
				class="input"
				bind:value={password}
				required
				placeholder="password"
			/>
		</div>
		<div
			class={state === State.REGISTER
				? 'transition-opacity duration-500 ease-in h-auto opacity-100 mb-4'
				: 'h-0 opacity-0'}
		>
			{#if state === State.REGISTER}
				<label for="name" class="block mb-2 text-sm font-medium text-white">Name</label>
				<input type="text" id="name" class="input" bind:value={name} required placeholder="name" />
			{/if}
		</div>

		<div class="flex gap-3 z-50">
			<button
				id="login"
				type={state === State.LOGIN ? 'submit' : 'reset'}
				on:click={() => handleStateChange(State.LOGIN)}
				class="button transition-all duration-300 ease-out {loginButtonWidth} truncate"
			>
				{#if state === State.LOGIN}
					Log In
				{:else}
					<LoginIcon />
				{/if}
			</button>

			<button
				id="register"
				type={state === State.REGISTER ? 'submit' : 'reset'}
				on:click={() => handleStateChange(State.REGISTER)}
				class="button transition-all duration-300 ease-out {registerButtonWidth} truncate"
			>
				{#if state === State.REGISTER}
					Register
				{:else}
					<RegisterIcon />
				{/if}
			</button>
		</div>
	</form>
</div>

<style lang="postcss">
	.input {
		@apply px-4 py-2 border border-gray-300 rounded-lg w-full;
	}

	.button {
		@apply bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded-lg mt-4;
	}
</style>
