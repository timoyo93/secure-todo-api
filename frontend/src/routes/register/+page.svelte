<script lang="ts">
	import { goto } from '$app/navigation';
	import { loginUser, registerUser } from '$root/lib/services/auth.service';
	import { Circle } from 'svelte-loading-spinners';

	let username: string;
	let password: string;
	let checkPassword: string;
	let loading: boolean = false;

	let formValid: boolean = true;
	let userAlreadyTaken: boolean = false;

	async function registerUserHandler() {
		loading = true;
		if (password !== checkPassword) {
			formValid = false;
			loading = false;
			return;
		}
		registerUser({ Username: username, Password: password })
			.then(() => {
				loginUser({ Username: username, Password: password }).then((res) => {
					console.log(res);
					loading = false;
					goto('/');
				});
			})
			.catch((err) => {
				if (err.response.status === 400) {
					userAlreadyTaken = true;
					loading = false;
					return;
				}
				loading = false;
				return;
			});
	}
</script>

<div class="flex flex-col items-center w-full mt-4 bg-white p-6 rounded-xl">
	<h1 class="text-xl">Please fill out the form below to register</h1>
	<form class="w-1/2 mt-5 flex flex-col" on:submit|preventDefault={registerUserHandler}>
		<input class="input" type="text" placeholder="Username*" required bind:value={username} />
		{#if userAlreadyTaken}
			<p class="text-center text-red-600 mb-4">User already taken. Try a different user</p>
		{/if}
		<input
			class="input"
			class:invalid={formValid === false}
			type="password"
			placeholder="Password*"
			required
			bind:value={password}
		/>
		<input
			class="input"
			class:invalid={formValid === false}
			type="password"
			placeholder="Repeat Password*"
			required
			bind:value={checkPassword}
		/>
		{#if !formValid}
			<p class="text-center text-red-600 mb-4">Passwords are not matching. Please check again</p>
		{/if}

		{#if loading}
			<div class="mx-auto">
				<Circle color="#3b82f6" size="30" />
			</div>
		{:else}
			<button
				class="bg-blue-500 text-white text-lg py-2 rounded-lg hover:shadow-md shadow-blue-700 hover:bg-blue-400 active:bg-blue-600 transition-colors duration-200"
				type="submit"
				>Register Now
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
