import type { AuthRequest, Todo } from '../models';
import axios from 'axios';

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL;

const client = axios.create({
	baseURL: import.meta.env.VITE_BACKEND_URL,
	headers: {
		'Content-Type': 'application/json'
	},
	withCredentials: true
});

export async function registerUser(request: AuthRequest) {
	return await apiCall(`${BACKEND_URL}/${Endpoints.AUTH_REGISTER}`, Method.POST, request, true);
}

export async function loginUser(request: AuthRequest) {
	// return await apiCall(`${BACKEND_URL}/${Endpoints.AUTH_LOGIN}`, Method.POST, request, true);
	const t = await client.post('/' + Endpoints.AUTH_LOGIN, request);
	return t;
}

export async function logoutUser() {
	return await apiCall(`${BACKEND_URL}/${Endpoints.AUTH_LOGOUT}`, Method.POST, {}, true);
}

export async function getTodos() {
	console.log(BACKEND_URL);
	return await apiCall(`${BACKEND_URL}/${Endpoints.API_TODOS}`, Method.GET, null);
}

export async function getTodo(id: string) {
	return await apiCall(`${BACKEND_URL}/${Endpoints.API_TODOS}/${id}`, Method.GET, {});
}

export async function addTodo(todo: Todo) {
	return await apiCall(`${BACKEND_URL}/${Endpoints.API_TODOS}`, Method.POST, todo);
}

export async function updateTodo(todo: Todo) {
	return await apiCall(`${BACKEND_URL}/${Endpoints.API_TODOS}`, Method.PUT, todo);
}

export async function deleteTodo(id: string) {
	return await apiCall(`${BACKEND_URL}/${Endpoints.API_TODOS}/${id}`, Method.DELETE, null);
}

function apiCall(url: string, method: Method, request: any, isAuth = false): Promise<Response> {
	return fetch(url, createRequest(method, request, isAuth));
}

function createRequest(method: Method, request: any, isAuth: boolean): RequestInit {
	return {
		method: method,
		// credentials: isAuth ? 'include' : 'same-origin',
		credentials: 'include',
		cache: 'no-cache',
		body: request ? JSON.stringify(request) : null,
		headers: {
			'Content-Type': 'application/json'
		}
	} as RequestInit;
}

enum Endpoints {
	AUTH = 'auth',
	AUTH_LOGIN = 'auth/login',
	AUTH_LOGOUT = 'auth/logout',
	AUTH_REGISTER = 'auth/register',
	API_TODO = 'api/v1/todo',
	API_TODOS = 'api/v1/todos'
}

enum Method {
	POST = 'post',
	GET = 'get',
	PUT = 'put',
	DELETE = 'delete'
}
