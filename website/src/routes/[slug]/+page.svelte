<script lang="ts">
	export let data;
	import { accessToken } from '$lib/auth';

	import { onMount } from 'svelte';

	let player;

	$: console.log('outside mount the store is:', $accessToken);

	onMount(() => {
		const script = document.createElement('script');
		script.src = 'https://sdk.scdn.co/spotify-player.js';
		script.async = true;

		document.body.appendChild(script);

		console.log('the access token from the on mount is:', $accessToken);

		window.onSpotifyWebPlaybackSDKReady = () => {
			player = new window.Spotify.Player({
				name: 'Web Playback SDK',
				getOAuthToken: (cb) => {
					cb(data.accessToken);
				},
				volume: 0.5
			});

			player.addListener('ready', ({ device_id }) => {
				console.log('Ready with Device ID', device_id);
			});

			player.addListener('not_ready', ({ device_id }) => {
				console.log('Device ID has gone offline', device_id);
			});

			player.connect();
		};
	});
</script>

<h1>{data.slug}</h1>
