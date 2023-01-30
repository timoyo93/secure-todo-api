import { writable, type Writable } from 'svelte/store';

export const isLoggedIn: Writable<boolean> = writable(false);
