<script>
	import LoginIcon from './LoginIcon.svelte';
	import RegisterIcon from './RegisterIcon.svelte';

	/**
	 * @enum { String }
	 */
	const State = {
		LOG: 'log',
		REGISTER: 'register'
	};

	/** @type {State}*/
	let state = State.LOG;
	let email = '';
	let password = '';
	let name = '';

	let loginButtonWidth = 'w-3/4';
	let registerButtonWidth = 'w-1/4';
	let loginDisplay = LoginIcon;
	let registerDisplay = RegisterIcon;

	/**
	 * Handles a login/register event
	 *
	 * @param e {SubmitEvent} - Form submission
	 * @listens
	 */
	const handleSubmit = (e) => {
		let action = e.submitter?.id;

		if (action === 'register') {
			console.log('registering...');
		} else {
			console.log('logging in...');
		}
	};

	/**
	 * Handles state change between registering or logging in
	 *
	 * 	@param new_state {State} - Register/Login button was click
	 *	@listens
	 */
	const handleStateChange = (new_state) => {
		if (new_state != state) {
			switch (new_state) {
				case State.LOG:
					state = State.LOG;
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
	<h3 class="text-center text-white mb-4 font-extrabold">klekfi</h3>

	<form on:submit|preventDefault={handleSubmit} autocomplete="off">
		{#if state === State.REGISTER}
			<div class="transition-all ease-in mb-4">
				<label for="name" class="block mb-2 text-sm font-medium text-white">Name</label>
				<input type="text" id="name" class="input" bind:value={name} required placeholder="name" />
			</div>
		{/if}
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

		<div class="flex gap-3">
			<button
				id="login"
				type={state === State.LOG ? 'submit' : 'button'}
				on:click={() => handleStateChange(State.LOG)}
				class="button transition-all duration-300 ease-in {loginButtonWidth} truncate"
			>
				{#if state === State.LOG}
					Log In
				{:else}
					<LoginIcon />
				{/if}
			</button>

			<button
				id="register"
				type={state === State.REGISTER ? 'submit' : 'button'}
				on:click={() => handleStateChange(State.REGISTER)}
				class="button transition-all duration-300 ease-in {registerButtonWidth} truncate"
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
