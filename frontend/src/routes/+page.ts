import { getTodos } from '$lib/api';
import type { PageLoad } from './$types';

export const load = (async () => {
	const res = await getTodos();
	console.log(res);
	if (res.ok) {
		return {
			todos: await res.json()
		};
	}
}) satisfies PageLoad;
