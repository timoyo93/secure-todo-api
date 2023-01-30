<script lang="ts">
	import { goto } from '$app/navigation';
	import TodoInput from '$lib/components/TodoInput.svelte';
	import TodoCard from '$lib/components/TodoCard.svelte';
	import { onMount } from 'svelte';
	import { isLoggedIn } from '../stores/auth.store';
	import type { PageData } from './$types';
	import { flip } from 'svelte/animate';
	import { fade, fly } from 'svelte/transition';

	export let data: PageData;

	onMount(() => {
		if (!$isLoggedIn) {
			goto('/login');
			return;
		}
	});
</script>

<svelte:head>
	<title>TODO App</title>
</svelte:head>

{#if $isLoggedIn}
	<div class="container" in:fade={{ duration: 300, delay: 300 }} out:fade={{ duration: 300 }}>
		<TodoInput />
		{#if data.todos}
			{#each data.todos as todo (todo.id)}
				<div
					animate:flip={{ duration: 500 }}
					in:fly={{ y: 200, duration: 500 }}
					out:fade={{ duration: 300 }}
				>
					<TodoCard {todo} />
				</div>
			{/each}
		{/if}
	</div>
{/if}
