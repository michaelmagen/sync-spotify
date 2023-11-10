<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import * as Alert from '$lib/components/ui/alert';
	import { AlertCircle } from 'lucide-svelte';
	import { PUBLIC_SPOTIFY_CLIENT_ID, PUBLIC_REDIRECT_URI } from '$env/static/public';
	import queryString from 'query-string';
	import { onMount } from 'svelte';
	import { page } from '$app/stores';

	let showFailedAuth = false;

	function getAuthURL() {
		const state = 'fddf';
		const scope =
			'streaming user-library-modify user-library-read user-read-email user-read-private';

		const url =
			'https://accounts.spotify.com/authorize?' +
			queryString.stringify({
				response_type: 'code',
				client_id: PUBLIC_SPOTIFY_CLIENT_ID,
				scope: scope,
				redirect_uri: PUBLIC_REDIRECT_URI,
				state: state
			});

		return url;
	}

	onMount(() => {
		// If failed to authroize, show error message
		const errorParam = $page.url.searchParams.get('error');
		if (errorParam === 'auth_failed') {
			showFailedAuth = true;
		}
	});
</script>

<div class="flex flex-col gap-5 justify-center items-center px-10 py-20">
	{#if showFailedAuth}
		<Alert.Root variant="destructive" class="max-w-sm">
			<AlertCircle class="h-4 w-4" />
			<Alert.Title>Failed to login!</Alert.Title>
			<Alert.Description>Please try again.</Alert.Description>
		</Alert.Root>
	{/if}
	<Button variant="default" size="lg" href={getAuthURL()}>Login With Spotify</Button>
</div>
