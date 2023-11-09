import { redirect } from '@sveltejs/kit';
import { accessToken, refreshToken } from '$lib/auth';
import { get } from 'svelte/store';

export const load = ({ request }) => {
	const storeAccessToken = get(accessToken);
	// If there already is a token in store, we know the user is logged in so stop here
	if (storeAccessToken) {
		console.log('the access token is already set');
		return;
	}
	// Get cookies from external API
	const externalCookies = request.headers.get('cookie')?.split('; ');
	const externalCookieMap = new Map<string, string>();
	// Create map of cookies
	if (externalCookies) {
		for (const cookie of externalCookies) {
			const [key, value] = cookie.split('=');
			externalCookieMap.set(key, value);
		}
	}
	// Get values from cookie map
	const newAccessToken = externalCookieMap.get('access_token');
	const newRefreshToken = externalCookieMap.get('refresh_token');
	const tokenExpiry = externalCookieMap.get('token_expiry');
	console.log('the token expiry:', tokenExpiry);
	// Set the svelte store vars
	if (newAccessToken) {
		accessToken.set(newAccessToken);
		console.log('the access Token is now set to:', get(accessToken));
	}

	if (newRefreshToken) {
		refreshToken.set(newRefreshToken);
		console.log('the refresh token is now set to:', get(refreshToken));
	}

	// If there is no access token in the store, redirect to login
	if (!get(accessToken)) {
		console.log('not authorized so redirecting to login');
		throw redirect(302, '/login');
	}

	return;
};
