<script lang="ts">
	import { goto, invalidateAll } from '$app/navigation';
	import TodoInput from '$lib/components/TodoInput.svelte';
	import TodoCard from '$lib/components/TodoCard.svelte';
	import { onMount } from 'svelte';
	import { isLoggedIn } from '../stores/auth.store';
	import type { PageData } from './$types';
	import { flip } from 'svelte/animate';
	import { fade, fly } from 'svelte/transition';
	import { checkAuth } from '$lib/api';

	export let data: PageData;

	async function authGuard() {
		const res = await checkAuth();
		console.log(res);
		if (res.ok) {
			$isLoggedIn = true;
			return;
		} else {
			$isLoggedIn = false;
			invalidateAll();
			goto('/login');
		}
	}

	onMount(() => {
		if (!$isLoggedIn) {
			goto('/login');
			return;
		} else {
			return;
		}
	});
</script>

<svelte:head>
	<title>TODO App</title>
</svelte:head>

{#await authGuard() then}
	<div class="container" in:fade={{ duration: 300, delay: 300 }} out:fade={{ duration: 300 }}>
		<TodoInput />
		{#if data.todos}
			<div class="todos-grid">
				{#each data.todos as todo (todo.id)}
					<div
						animate:flip={{ duration: 500 }}
						in:fly={{ y: 200, duration: 500 }}
						out:fade={{ duration: 300 }}
					>
						<TodoCard {todo} />
					</div>
				{/each}
			</div>
		{/if}
	</div>
{/await}

<style>
	.todos-grid {
		display: grid;
		gap: 1rem;
		justify-content: center;
		grid-template-columns: repeat(auto-fill, minmax(300px, 450px));
	}
</style>
