<script lang="ts">
	import TodosView from '$root/lib/components/TodosView.svelte';
	import { goto } from '$app/navigation';
	import { checkAuth } from '$root/lib/services/auth.service';
	import { authStore } from '$root/stores/store';
	import { onMount } from 'svelte';
	onMount(async () => {
		checkAuth()
			.then(() => {
				authStore.set(true);
			})
			.catch((err) => {
				if (err.response.status === 401) {
					authStore.set(false);
					goto('/login');
				}
			});
	});
</script>

<div class="w-full">
	{#if $authStore}
		<TodosView />
	{:else}
		<p>You must login first</p>
	{/if}
</div>

<style lang="scss">
</style>
