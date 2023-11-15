<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { accessToken } from '$lib/auth';
	import { onMount } from 'svelte';
	import { PUBLIC_API_URL } from '$env/static/public';
	import { goto } from '$app/navigation';

	//$: console.log('from base page access token is:', $accessToken);
	// TODO: The above does not work because :
	/// Stores do not magically get transferred between server and client, you should not use them on the server anyway. There also would just be one instance for all users. (If you need server-side state you have to make sure to separate it by user/session.) Load the data from the server e.g. via a layout load function, that way all pages using said layout will have access via the data property.
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
<a href="/login">login</a>

<Button>Create New Room</Button>
<Button
	on:click={() => {
		goto('/fjkdsla');
	}}>Join a room</Button
>
