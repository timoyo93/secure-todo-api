<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { deleteTodo, updateTodo } from '$lib/api';
	import type { Todo } from '../../models';

	export let todo: Todo;
	let showError = false;
	let errorMessage = '';

	async function markDone() {
		todo.completed = !todo.completed;
		showError = false;
		const res = await updateTodo(todo);
		if (res.ok) {
			return;
		} else {
			showError = true;
			errorMessage = 'Error updating Todo, try again';
			todo.completed = !todo.completed;
		}
	}

	async function removeTodo() {
		const res = await deleteTodo(todo.id);
		if (res.ok) {
			invalidateAll();
			return;
		} else {
			showError = true;
			errorMessage = 'Error deleting todo, try again';
			return;
		}
	}
</script>

<section>
	<div class="grid">
		<article>
			<h2 class="text-center">
				{#if todo.completed}
					<del>{todo.name}</del>
				{:else}
					{todo.name}
				{/if}
			</h2>
			{#if showError}
				<p class="text-center">{errorMessage}</p>
			{/if}
			<footer class="flex-actions">
				<button on:click={markDone} class="outline width--half">Mark Done</button>
				<button on:click={removeTodo} class="outline contrast width--half">Delete</button>
			</footer>
		</article>
	</div>
</section>

<style>
	.text-center {
		text-align: center;
	}
	.width--half {
		width: 50%;
	}
	.flex-actions {
		display: flex;
		gap: 1rem;
	}
</style>
