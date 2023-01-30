<script lang="ts">
	import { fade } from 'svelte/transition';
	let loading = false;
	let userExisting = false;
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
			<form on:submit|preventDefault>
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
				<input
					type="password"
					id="confirm-password"
					name="confirm-password"
					placeholder="Confirm Password"
					aria-label="Confirm Password"
					required
				/>
				{#if !loading}
					<button type="submit">Register</button>
				{:else}
					<button aria-busy="true" />
				{/if}
				{#if userExisting}
					<p class="error">User with this Username already exists, try another name</p>
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
