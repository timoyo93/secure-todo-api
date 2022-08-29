import API from '$root/lib/api';
import type { AxiosError, AxiosResponse } from 'axios';
import { env } from '$env/dynamic/public';

const client = new API(env.PUBLIC_BACKEND_URL);

export interface Todo {
	id: string;
	name: string;
	completed: boolean;
}

export const addTodo = async (todo: Todo): Promise<AxiosResponse<Todo | string> | AxiosError> => {
	const response = await client.post<Todo | string>('/api/v1/todo', todo);
	return response;
};

export const removeTodo = async (todoId: string): Promise<AxiosResponse<string> | AxiosError> => {
	const response = await client.delete<string>(`/api/v1/todo/${todoId}`, '');
	return response;
};

export const getTodos = async (): Promise<AxiosResponse<Todo[] | string> | AxiosError> => {
	const response = await client.get<Todo[] | string>('/api/v1/todos', '');
	return response;
};

export const updateTodo = async (
	todo: Todo
): Promise<AxiosResponse<Todo | string> | AxiosError> => {
	const response = await client.put<Todo | string>(`/api/v1/todo/${todo.id}`, todo);
	return response;
};
