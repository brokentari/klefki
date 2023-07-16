<script>
	/** @type {import('$lib/types').User}*/
	export let current_user;
	/** @type boolean*/
	export let loggedIn;

	const handleLogout = () => {
		const storedUser = localStorage.getItem('userId');
		if (storedUser) {
			localStorage.removeItem('userId');
			loggedIn = false;
		}
	};
</script>

<nav class="bg-gray-800 fixed w-full top-0 left-0 right-0">
	<div class="mx-auto px-4 sm:px-6 lg:px-8">
		<div class="flex justify-between h-16">
			<!-- Logo -->
			<div class="flex-shrink-0 flex items-center">
				<a href="/" class="text-white text-lg font-semibold">klekfi</a>
			</div>

			<!-- Navigation Links -->
			<div class="flex items-center">
				<div class="hidden md:block">
					<a
						href="google.com"
						class="text-gray-300 font-bold hover:bg-gray-700 px-3 py-2 rounded-md text-sm"
						>{current_user ? current_user.name : 'undefined'}</a
					>
				</div>
				<div class="hidden md:block">
					<button on:click={handleLogout}>
						<a
							href="/"
							class="text-gray-300 hover:bg-gray-700 px-3 py-2 rounded-md text-sm font-medium"
							>log out
						</a>
					</button>
				</div>
			</div>
		</div>
	</div>
</nav>
