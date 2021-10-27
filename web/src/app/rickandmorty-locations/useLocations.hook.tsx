
import { LocationType } from './types';
import { useEffect, useState } from 'react';
import { LocationsService } from './service'

export function useLocations(): LocationType[] {
    const [state, setState] = useState<LocationType[]>([]);

    useEffect(() => {
        async function fetchLocations() {
            try {
                const response = await LocationsService.get();

                setState(response);

            } catch (error) {
                console.log(error);
            }
        }

        fetchLocations();
    }, []);

    return state;
}
