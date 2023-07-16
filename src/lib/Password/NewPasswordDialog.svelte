<div>
	<div
		class="fixed z-10 inset-0 flex items-center justify-center bg-black bg-opacity-50"
		id="overlay"
	/>
	<div class="fixed z-10 inset-0 overflow-y-auto backdrop:bg-gray-500">
		<div class="flex items-center justify-center min-h-screen p-4 text-center">
			<div class="bg-white rounded-lg shadow-lg p-8">
				<h2 class="text-xl font-semibold mb-4">{selected_password.name}</h2>
				<p>{selected_password.username}</p>
				<button
					on:click={handleCloseDialog}
					class="mt-4 bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
					id="closeBtn">close</button
				>
			</div>
		</div>
	</div>
</div>
