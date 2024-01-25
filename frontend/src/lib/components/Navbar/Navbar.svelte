<script>
	import { page } from '$app/stores';
	import { Navbar, NavBrand, NavLi, NavUl, Button, P } from 'flowbite-svelte';
	import { DarkMode } from 'flowbite-svelte';

	$: activeUrl = $page.url.pathname;
</script>

<Navbar color="primary">
	<NavBrand href="/">
		<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">Navbar</span>
	</NavBrand>
	<div class="flex gap-2 md:order-2">
		<DarkMode />
		{#if !$page.data.token}
			<Button size="sm" href="/auth/login">Login</Button>
		{:else}
			<form action="/auth/logout" method="POST">
				<Button color="red" type="submit">Signout</Button>
			</form>
		{/if}
	</div>
	<NavUl {activeUrl}>
		<NavLi href="/">Home</NavLi>

		<NavLi href="/courses">Courses</NavLi>
	</NavUl>
</Navbar>
