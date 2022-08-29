<script lang="ts">
	import { authStore } from '$root/stores/store';
	import { logoutUser } from '$root/lib/services/auth.service';
	import { goto } from '$app/navigation';
	export let logoTitle = 'LOGO';

	async function logoutHandler() {
		try {
			const res = await logoutUser();
			if (res.status === 200) {
				authStore.set(false);
				goto('/login');
			}
		} catch (err) {
			console.log(err);
		}
	}
</script>

<nav class="bg-white w-full shadow-md">
	<div class="flex justify-between items-center align-middle py-3 w-3/5 container mx-auto">
		<div class="text-xl"><a href="/">{logoTitle}</a></div>
		<div class="flex space-x-4">
			{#if $authStore}
				<button class="button bg-red-500 text-white hover:bg-red-400" on:click={logoutHandler}
					>Log Out</button
				>
			{:else}
				<button class="button bg-green-600 hover:bg-green-400"><a href="/login">Login</a></button>
				<button class="button bg-blue-500 hover:bg-blue-400"
					><a href="/register">Register</a></button
				>
			{/if}
		</div>
	</div>
</nav>

<style lang="scss">
	.button {
		@apply hover:shadow-lg rounded-lg px-4 py-2 text-white transition-colors;
	}
</style>
