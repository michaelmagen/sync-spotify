// TODO: Move this to top level of layout, send the access token down through data, do not set it in server. In any page, if there is not a cookie with the access code, then send to login. Make the cookie not expire so fast.
export const load = ({ locals }) => {
	return {
		accessToken: locals.accessToken
	};
};
