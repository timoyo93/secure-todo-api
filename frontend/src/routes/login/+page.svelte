<script lang="ts">
	import { goto } from '$app/navigation';
	import { loginUser } from '$root/lib/services/auth.service';
	import { authStore } from '$root/stores/store';
	import { Circle } from 'svelte-loading-spinners';

	let username: string;
	let password: string;
	let loading: boolean = false;
	let formValid: boolean = true;
	let errorMessage: string = '';

	async function loginUserHandler() {
		loading = true;
		formValid = true;
		loginUser({ Username: username, Password: password })
			.then((res) => {
				console.log(res);
				loading = false;
				authStore.set(true);
				goto('/');
			})
			.catch((err) => {
				loading = false;
				formValid = false;
				if (err.response.status === 400) {
					errorMessage = 'Username or password are wrong, please try again';
				} else if (err.response.status === 404) {
					errorMessage = 'You are not registered yet. Please register first!';
				} else {
					errorMessage = 'An unexpected error occured';
				}
			});
	}
</script>

<div class="flex flex-col items-center w-full mt-4 bg-white p-6 rounded-xl">
	<h1 class="text-xl">Please log in</h1>
	<form class="w-1/2 mt-5 flex flex-col" on:submit|preventDefault={loginUserHandler}>
		<input
			class="input"
			class:invalid={formValid === false}
			type="text"
			placeholder="Username*"
			required
			bind:value={username}
		/>
		<input
			class="input"
			class:invalid={formValid === false}
			type="password"
			placeholder="Password*"
			required
			bind:value={password}
		/>
		{#if !formValid}
			<p class="text-center text-red-600 mb-4">{errorMessage}</p>
		{/if}

		{#if loading}
			<div class="mx-auto">
				<Circle color="#3b82f6" size="30" />
			</div>
		{:else}
			<button
				class="bg-blue-500 text-white text-lg py-2 rounded-lg hover:shadow-md shadow-blue-700 hover:bg-blue-400 active:bg-blue-600 transition-colors duration-200"
				type="submit"
				>Login
			</button>
		{/if}
	</form>
</div>

<style lang="scss">
	.input {
		@apply w-full border-2 rounded-md p-2 shadow-sm focus:border-blue-500 focus:ring-blue-500 mb-4;
	}
	.invalid {
		@apply border-red-600;
	}
</style>
