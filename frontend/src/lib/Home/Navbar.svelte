<script>
	/** @type {import('../../routes/+page.svelte').User}*/
	export let current_user;
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
						class="text-gray-300 hover:bg-gray-700 px-3 py-2 rounded-md text-sm font-medium"
						>{current_user.name}</a
					>
				</div>
			</div>
		</div>
	</div>
</nav>
