<div class="fixed inset-0 flex items-center justify-center bg-gray-900 bg-opacity-75">
	<div class="bg-white rounded-lg p-8 shadow-md">
		<h2 class="text-2xl font-bold mb-6">Form Dialog</h2>
		<div class="field">
			<label for="website" class="font-semibold">website: </label>
			<input type="text" id="website" class="border border-gray-300 rounded-md p-2" />
		</div>

		<div class="field">
			<label for="website" class="font-semibold">name: </label>
			<input type="text" id="website" class="border border-gray-300 rounded-md p-2" />
		</div>

		<div class="field">
			<label for="website" class="font-semibold">username: </label>
			<input type="text" id="website" class="border border-gray-300 rounded-md p-2" />
		</div>

		<div class="field">
			<label for="website" class="font-semibold">password: </label>
			<input type="text" id="website" class="border border-gray-300 rounded-md p-2" />
		</div>
	</div>
</div>

<style lang="postcss">
	.field { 
		display: grid;
		gap: 1rem;
		grid-template-columns: repeat(2, 1fr);
	}
</style>
