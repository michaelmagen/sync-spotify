import { redirect } from '@sveltejs/kit';
import { accessToken } from '$lib/auth';
import { get } from 'svelte/store';

export const load = ({ cookies }) => {
	// TODO: Make sure this function works properly once everything is set up
	// This may be outdated, since the server automatically refreshes for you, then deletes the old access token.
	console.log('the cookeies has', cookies.getAll());

	const storeAccessToken = get(accessToken);
	// If there already is a token in store, we know the user is logged in so stop here
	if (storeAccessToken) {
		console.log('the access token is already set');
		return;
	}

	// Get values from cookie map
	const newAccessToken = cookies.get('access_token');
	// Set the svelte store vars
	if (newAccessToken) {
		accessToken.set(newAccessToken);
		console.log('the access Token is now set to:', get(accessToken));
	}

	// If there is no access token in the store, redirect to login
	if (!get(accessToken)) {
		console.log('not authorized so redirecting to login');
		throw redirect(302, '/login');
	}

	return;
};
