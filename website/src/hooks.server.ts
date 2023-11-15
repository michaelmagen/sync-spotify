import { redirect, type Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	// Get accessToken from cookie
	const accessToken = event.cookies.get('access_token');
	event.locals.accessToken = event.cookies.get('access_token');

	console.log(event.url.pathname, accessToken);

	// Redirect to /login if user without accessToken tries to access any other page
	if (!(event.url.pathname === '/login') && !accessToken) {
		throw redirect(302, '/login');
	}

	const response = await resolve(event);
	return response;
};
