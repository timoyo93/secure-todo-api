import { goto } from '$app/navigation';
import { Endpoints } from '$lib/api';
import { redirect } from '@sveltejs/kit';
import { isLoggedIn } from '../stores/auth.store';
import type { PageLoad } from './$types';

export const load = (async ({ fetch }) => {
	const auth = await fetch(`${import.meta.env.VITE_BACKEND_URL}/${Endpoints.AUTH}`, {
		credentials: 'include'
	});
	if (auth.status === 401) {
		isLoggedIn.set(false);
		throw redirect(301, '/login');
	}
	if (auth.status === 200) {
		isLoggedIn.set(true);
		const res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/${Endpoints.API_TODOS}`, {
			credentials: 'include'
		});
		if (res.ok) {
			return {
				todos: await res.json()
			};
		}
	}
}) satisfies PageLoad;
