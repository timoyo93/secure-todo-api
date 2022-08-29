/* eslint-disable @typescript-eslint/no-explicit-any */
import type { AxiosError, AxiosInstance, AxiosResponse, Method } from 'axios';
import axios from 'axios';

export default class API {
	_client: AxiosInstance;

	constructor(baseUrl: string) {
		this._client = axios.create({
			baseURL: baseUrl,
			withCredentials: true
		});
	}

	async apiRequest<Response>(
		method: Method,
		url: string,
		request?: any
	): Promise<AxiosResponse<Response> | AxiosError> {
		const res = await this._client({
			method: method,
			url: url,
			data: request
		});
		return res;
	}

	get<T>(url: string, request: any) {
		return this.apiRequest<T>('get', url, request);
	}
	delete<T>(url: string, request: any) {
		return this.apiRequest<T>('delete', url, request);
	}
	post<T>(url: string, request: any) {
		return this.apiRequest<T>('post', url, request);
	}
	put<T>(url: string, request: any) {
		return this.apiRequest<T>('put', url, request);
	}
	patch<T>(url: string, request: any) {
		return this.apiRequest<T>('patch', url, request);
	}
}
