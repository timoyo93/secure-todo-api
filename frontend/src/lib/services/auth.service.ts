import API from '$root/lib/api';
import type { AxiosError, AxiosResponse } from 'axios';
import { env } from '$env/dynamic/public';

const client = new API(env.PUBLIC_BACKEND_URL);

interface AuthRequest {
	Username: string;
	Password: string;
}

interface User {
	Id: string;
	Token: string;
	Username: string;
	Hash: string;
}

export const registerUser = async (
	req: AuthRequest
): Promise<AxiosResponse<string> | AxiosError> => {
	const response = await client.post<string>('/auth/register', req);
	return response;
};

export const loginUser = async (
	req: AuthRequest
): Promise<AxiosResponse<User | string> | AxiosError> => {
	const response = await client.post<User | string>('/auth/login', req);
	return response;
};

export const logoutUser = async (): Promise<AxiosResponse<string> | AxiosError> => {
	const response = await client.post<string>('/auth/logout', null);
	return response;
};

export const checkAuth = async (): Promise<AxiosResponse<string> | AxiosError> => {
	const response = await client.get<string>('/auth', null);
	return response;
};
