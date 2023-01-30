<script lang="ts">
	import { goto } from '$app/navigation';
	import { logoutUser } from '$lib/api';
	import '@picocss/pico';
	import { isLoggedIn } from '../stores/auth.store';
	import '../app.css';

	async function logout() {
		const res = await logoutUser();
		if (res.ok) {
			$isLoggedIn = false;
		}
		goto('/login');
		return;
	}
</script>

<nav class="container-fluid">
	<ul>
		<li><a href="/" class="contrast"><strong>TODO App</strong></a></li>
	</ul>
	<ul>
		{#if $isLoggedIn}
			<li><button class="primary outline" on:click={logout}>Logout</button></li>
		{:else}
			<li><a href="/login" class="primary">Login</a></li>
			<li><a href="/register" class="contrast">Register</a></li>
		{/if}
	</ul>
</nav>
<slot />
