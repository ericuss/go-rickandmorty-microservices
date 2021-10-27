import { FC } from "react";
import { useCharacters } from "./useCharacters.hook";
import { CharacterType } from "./types";
import './index.css';

type CharacterTypeProps = {
    character: CharacterType
}

export const Characters: FC = () => {
    const characters = useCharacters();

    const Character: FC<CharacterTypeProps> = ({ character }) => {
        return (
            <div className="character">
                <img src={character.Image} />
                <div className="character__name">{character.Name}</div>
            </div>
        );
    };

    return (
        <div className="characters">
            {characters.map((c) => <Character character={c} />)}
        </div>
    );

}