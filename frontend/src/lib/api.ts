import type { AuthRequest, Todo } from '../models';

const BACKEND_URL = import.meta.env.VITE_BACKEND_URL;

export async function checkAuth() {
	return await apiCall(`${BACKEND_URL}/${Endpoints.AUTH}`, Method.GET, null);
}

export async function registerUser(request: AuthRequest) {
	return await apiCall(`${BACKEND_URL}/${Endpoints.AUTH_REGISTER}`, Method.POST, request);
}

export async function loginUser(request: AuthRequest) {
	return await apiCall(`${BACKEND_URL}/${Endpoints.AUTH_LOGIN}`, Method.POST, request);
}

export async function logoutUser() {
	return await apiCall(`${BACKEND_URL}/${Endpoints.AUTH_LOGOUT}`, Method.POST, {});
}

export async function getTodos() {
	return await apiCall(`${BACKEND_URL}/${Endpoints.API_TODOS}`, Method.GET, null);
}

export async function getTodo(id: string) {
	return await apiCall(`${BACKEND_URL}/${Endpoints.API_TODO}/${id}`, Method.GET, {});
}

export async function addTodo(todo: Todo) {
	return await apiCall(`${BACKEND_URL}/${Endpoints.API_TODO}`, Method.POST, todo);
}

export async function updateTodo(todo: Todo) {
	return await apiCall(`${BACKEND_URL}/${Endpoints.API_TODO}`, Method.PUT, todo);
}

export async function deleteTodo(id: string) {
	return await apiCall(`${BACKEND_URL}/${Endpoints.API_TODO}/${id}`, Method.DELETE, null);
}

function apiCall(url: string, method: Method, request: any): Promise<Response> {
	return fetch(url, createRequest(method, request));
}

function createRequest(method: Method, request: any): RequestInit {
	return {
		method: method,
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
