<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import {
		PUBLIC_API_URL,
		PUBLIC_SPOTIFY_CLIENT_ID,
		PUBLIC_REDIRECT_URI
	} from '$env/static/public';
	import queryString from 'query-string';

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
</script>

<div class="flex justify-center items-center px-10 py-20">
	<Button variant="default" size="lg" href={getAuthURL()}>Login With Spotify</Button>
</div>
