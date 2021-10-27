
import axios, { AxiosResponse } from 'axios';

export const Instance = axios.create({
	baseURL: 'http://host.docker.internal:5008/api/',
	timeout: 15000,
	headers: {
	  'Content-Type': 'application/json',
	  'X-Requested-With': 'XMLHttpRequest'
	}
});

export const responseBody = (response: AxiosResponse) => response.data;

export const Requests = {
	get: (url: string) => Instance.get(url).then(responseBody),
};
