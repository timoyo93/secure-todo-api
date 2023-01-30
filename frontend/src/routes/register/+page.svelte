<script lang="ts">
	import { goto } from '$app/navigation';
	import { loginUser, registerUser } from '$lib/api';
	import { fade } from 'svelte/transition';
	import type { AuthRequest } from '../../models';
	import { isLoggedIn } from '../../stores/auth.store';
	let loading = false;
	let password = '';
	let checkPassword = '';
	let arePasswordEqual = false;
	let showError = false;
	let errorMessage = '';

	function checkIfPasswordIsEqual() {
		if (password.length > 0) {
			const isEqual = password === checkPassword;
			arePasswordEqual = isEqual;
		} else {
			return;
		}
	}

	async function register(e: SubmitEvent) {
		loading = true;
		showError = false;
		const formData = new FormData(e.target as HTMLFormElement);
		const request: AuthRequest = {
			username: String(formData.get('username')),
			password: checkPassword
		};
		const registerResponse = await registerUser(request);
		if (!registerResponse.ok) {
			const error = await registerResponse.json();
			showError = true;
			loading = false;
			errorMessage = `${error}, try again`;
			return;
		}
		const authResponse = await loginUser(request);
		if (authResponse.ok) {
			$isLoggedIn = true;
			loading = false;
			goto('/');
			return;
		} else {
			showError = true;
			loading = false;
			errorMessage = 'Error during login, try<a href="/login">again</a> ';
		}
	}
</script>

<svelte:head>
	<title>Register</title>
</svelte:head>

<div class="page-center" in:fade={{ duration: 500, delay: 300 }}>
	<div class="container">
		<article>
			<hgroup>
				<h2>Register</h2>
				<h3>Already registered? <a href="/login">Log in</a></h3>
			</hgroup>
			<form on:input={checkIfPasswordIsEqual} on:submit|preventDefault={register}>
				<input
					type="text"
					id="username"
					name="username"
					placeholder="Username"
					aria-label="Username"
					required
				/>
				<input
					bind:value={password}
					type="password"
					id="password"
					name="password"
					placeholder="Password"
					aria-label="Password"
					required
				/>
				<input
					bind:value={checkPassword}
					type="password"
					id="confirm-password"
					name="confirm-password"
					placeholder="Confirm Password"
					aria-label="Confirm Password"
					required
				/>
				{#if !loading}
					<button disabled={!arePasswordEqual} type="submit">Register</button>
				{:else}
					<button aria-busy="true" />
				{/if}
				{#if showError}
					<p class="error">{@html errorMessage}</p>
				{/if}
			</form>
		</article>
	</div>
</div>

<style>
	.page-center {
		margin-top: 12rem;
		display: flex;
		flex-direction: column;
		justify-content: center;
	}
	.error {
		text-align: center;
		color: red;
	}
</style>
