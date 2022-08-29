import { writable, type Writable } from 'svelte/store';

export const authStore: Writable<boolean> = writable(false);
