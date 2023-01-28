<script lang="ts">
	import { getTodos, loginUser } from '$lib/api';
	import { onMount } from 'svelte';

	onMount(async () => {
		const response = await loginUser({ username: 'timo', password: 'asdf' });
		console.log(response);
		const todos = await getTodos();
		console.log(todos);
	});
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
