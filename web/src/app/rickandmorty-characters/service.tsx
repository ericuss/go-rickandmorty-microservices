
import axios, { AxiosResponse } from 'axios';
import { CharacterType } from './types';

const instance = axios.create({
	baseURL: 'http://host.docker.internal:5008/api/',
	timeout: 15000,
	headers: {
	  'Content-Type': 'application/json',
	  'X-Requested-With': 'XMLHttpRequest'
	}
});

const responseBody = (response: AxiosResponse) => response.data;

const requests = {
	get: (url: string) => instance.get(url).then(responseBody),
};

export const CharactersService = {
	get: (): Promise<CharacterType[]> => requests.get('characters'),
};
