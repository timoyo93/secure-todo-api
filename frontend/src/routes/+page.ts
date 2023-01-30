import { Endpoints } from '$lib/api';
import { redirect } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load = (async ({ fetch }) => {
	try {
		const auth = await fetch(`${import.meta.env.VITE_BACKEND_URL}/${Endpoints.AUTH}`, {
			credentials: 'include'
		});
		if (!auth.ok) {
			throw redirect(301, '/login');
		}
		const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/${Endpoints.API_TODOS}`, {
			credentials: 'include'
		});
		if (res.ok) {
			return {
				todos: await res.json()
			};
		}
	} catch {
		throw redirect(301, '/login');
	}
}) satisfies PageLoad;
