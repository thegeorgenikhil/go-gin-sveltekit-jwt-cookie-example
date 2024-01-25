<script lang="ts">
	import Navbar from '$lib/components/Navbar/Navbar.svelte';
	import { Card, Button, Label, Input } from 'flowbite-svelte';
	import { EyeOutline, EyeSlashOutline } from 'flowbite-svelte-icons';

	let show = false;
</script>

<Navbar />

<Card class="mx-auto mt-20">
	<form class="flex flex-col space-y-6" method="POST" action="?/signup">
		<h3 class="text-xl font-medium text-gray-900 dark:text-white">Sign up to our platform</h3>
		<Label class="space-y-2">
			<span>Email</span>
			<Input type="email" name="email" placeholder="johndoe@email.com" required />
		</Label>
		<Label class="space-y-2">
			<span>Password</span>
			<Input name="password" type={show ? 'text' : 'password'} placeholder="•••••" required>
				<button
					slot="left"
					on:click|preventDefault={() => (show = !show)}
					class="pointer-events-auto"
				>
					{#if show}
						<EyeOutline class="h-6 w-6" />
					{:else}
						<EyeSlashOutline class="h-6 w-6" />
					{/if}
				</button>
			</Input>
		</Label>
		<Button type="submit" class="w-full">Signup</Button>
		<p class="text-center text-sm font-medium text-gray-500 dark:text-gray-300">
			Already have an account? <a
				href="/auth/signup"
				class="text-center text-primary-700 hover:underline dark:text-primary-500"
			>
				Sign in
			</a>
		</p>
	</form>
</Card>
