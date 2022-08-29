<script lang="ts">
	import { onMount } from 'svelte';
	import { addTodo, getTodos, removeTodo, updateTodo, type Todo } from '../services/todo.service';
	import type { AxiosResponse, AxiosError } from 'axios';
	import TodoCard from './TodoCard.svelte';

	let newTodo: string = '';

	let todos: Todo[] = [];

	onMount(async () => {
		getTodos()
			.then((res) => {
				const response = res as AxiosResponse<Todo[]>;
				todos = response.data !== null ? response.data : [];
			})
			.catch();
	});

	function addNewTodo(): void {
		const todo = { name: newTodo, completed: false, id: '' };
		addTodo(todo)
			.then(() => (newTodo = ''))
			.finally(() => {
				getTodos()
					.then((res) => {
						const response = res as AxiosResponse<Todo[]>;
						todos = response.data !== null ? response.data : [];
					})
					.catch();
			});
	}
	function deleteTodo(todo: Todo) {
		removeTodo(todo.id).then(() => {
			getTodos()
				.then((res) => {
					const response = res as AxiosResponse<Todo[]>;
					todos = response.data !== null ? response.data : [];
				})
				.catch();
		});
	}

	function markTodoAsCompleted(todo: Todo) {
		todo.completed = true;
		updateTodo(todo).then(() => {
			getTodos()
				.then((res) => {
					const response = res as AxiosResponse<Todo[]>;
					todos = response.data !== null ? response.data : [];
				})
				.catch();
		});
	}
</script>

<div class="flex flex-col items-center mt-4">
	<form class="mb-2" on:submit|preventDefault={addNewTodo}>
		<input
			class="w-96"
			type="text"
			name="newTodo"
			id="newTodo"
			placeholder="Enter new Todo"
			autocomplete="off"
			bind:value={newTodo}
		/>
	</form>
	<div>
		{#each todos as todo (todo.id)}
			<TodoCard {todo} {deleteTodo} updateTodo={markTodoAsCompleted} />
		{/each}
	</div>
</div>

<style lang="scss">
</style>
