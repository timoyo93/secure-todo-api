<script lang="ts">
	import type { Todo } from '$root/lib/services/todo.service';
	export let todo: Todo;
	export let deleteTodo = (todo: Todo): void => {};
	export let updateTodo = (todo: Todo): void => {};
</script>

<div class="flex items-center justify-between bg-white shadow-md p-4 w-96 mb-2">
	<p class:completed={todo.completed === true}>{todo.name}</p>
	<div class="flex align-middle space-x-2">
		{#if !todo.completed}
			<button class="bg-green-500 text-white p-2" on:click={() => updateTodo(todo)}>DONE</button>
		{/if}
		<button class="bg-red-600 text-white p-2" on:click={() => deleteTodo(todo)}>X</button>
	</div>
</div>

<style lang="scss">
	.completed {
		@apply line-through;
	}
</style>
