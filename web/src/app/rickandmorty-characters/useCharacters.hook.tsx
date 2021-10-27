
import { CharacterType } from './types';
import { useEffect, useState } from 'react';
import { CharactersService } from './service'

export function useCharacters(): CharacterType[] {
    const [state, setState] = useState<CharacterType[]>([]);

    useEffect(() => {
        async function fetchCharacters() {
            try {
                const response = await CharactersService.get();

                setState(response);

            } catch (error) {
                console.log(error);
            }
        }

        fetchCharacters();
    }, []);

    return state;
}
