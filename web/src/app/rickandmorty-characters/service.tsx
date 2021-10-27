
import { CharacterType } from './types';
import { Requests } from '../core/http/serviceCore'

export const CharactersService = {
	get: (): Promise<CharacterType[]> => Requests.get('characters'),
};
