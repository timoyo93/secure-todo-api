<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { addTodo } from '$lib/api';
	import type { Todo } from '../../models';

	async function createTodo(e: SubmitEvent) {
		const formElement = e.target as HTMLFormElement;
		const formData = new FormData(formElement);
		const name = formData.get('todo-name');
		const todo: Todo = {
			id: '',
			completed: false,
			name: String(name)
		};
		const res = await addTodo(todo);
		if (res.ok) {
			invalidateAll();
			formElement.reset();
		}
	}
</script>

<form on:submit|preventDefault={createTodo} class="container flex-form">
	<input
		type="text"
		name="todo-name"
		id="todo-name"
		placeholder="Add Todo..."
		aria-label="Todo Name"
		autocomplete="off"
		required
	/>
	<button type="submit">Add Todo</button>
</form>

<style>
	.flex-form {
		margin-top: 2rem;
		display: flex;
		flex-wrap: wrap;
		gap: 1rem;
	}
	.flex-form > * {
		flex-grow: 1;
		flex-shrink: 1;
		flex-basis: content;
	}
	.flex-form > button {
		width: 30%;
	}
	.flex-form > input {
		width: 70%;
	}
</style>
