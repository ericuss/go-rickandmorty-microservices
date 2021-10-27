
import { LocationType } from './types';
import { Requests } from '../core/http/serviceCore'

export const LocationsService = {
	get: (): Promise<LocationType[]> => Requests.get('locations'),
};
