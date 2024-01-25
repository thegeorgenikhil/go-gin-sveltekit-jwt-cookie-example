<script lang="ts">
	import Navbar from '$lib/components/Navbar/Navbar.svelte';
	import { Card, Heading } from 'flowbite-svelte';
	import { page } from '$app/stores';
</script>

<Navbar />

<Card class="mx-auto mt-20">
	<Heading class="text-center">
		Your id is {$page.data.id}
	</Heading>
</Card>
