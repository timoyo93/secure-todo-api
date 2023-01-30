<script lang="ts">
	import { goto } from '$app/navigation';
	import { loginUser } from '$lib/api';
	import type { AuthRequest } from '../../models';
	import { isLoggedIn } from '../../stores/auth.store';
	import { fade } from 'svelte/transition';

	let loading = false;

	async function login(e: SubmitEvent) {
		loading = true;
		const formData = new FormData(e.target as HTMLFormElement);
		const authRequest: AuthRequest = {
			username: String(formData.get('username')),
			password: String(formData.get('password'))
		};
		const response = await loginUser(authRequest);
		if (response.ok) {
			loading = false;
			$isLoggedIn = true;
			goto('/');
			return;
		}
		loading = false;
	}
</script>

<svelte:head>
	<title>Login</title>
</svelte:head>

<div class="page-center" in:fade={{ duration: 500, delay: 300 }}>
	<div class="container">
		<article>
			<hgroup>
				<h2>Login</h2>
				<h3>Not registered yet? <a href="/register">Register here</a></h3>
			</hgroup>
			<form on:submit|preventDefault={login}>
				<input
					type="text"
					id="username"
					name="username"
					placeholder="Username"
					aria-label="Username"
					required
				/>
				<input
					type="password"
					id="password"
					name="password"
					placeholder="Password"
					aria-label="Password"
					required
				/>
				{#if !loading}
					<button type="submit">Login</button>
				{:else}
					<button aria-busy="true" />
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
</style>
