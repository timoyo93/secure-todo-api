<script lang="ts">
	import Navbar from '$root/lib/components/navbar.svelte';
	import '$root/app.css';
</script>

<div class="bg-gray-200 h-screen">
	<Navbar logoTitle="Awesome TODOs" />
	<div class="container w-3/5 mx-auto flex justify-center">
		<slot />
	</div>
</div>

<style lang="scss">
</style>
