import { FC } from "react";
import { useLocations } from "./useLocations.hook";
import { LocationType } from "./types";
import './index.css';

type LocationTypeProps = {
    location: LocationType
}

export const Locations: FC = () => {
    const locations = useLocations();

    const Location: FC<LocationTypeProps> = ({ location }) => {
        return (
            <div className="location">
                <div className="location__name">{location.Name}</div>
                <div className="location__dimension">{location.Dimension}</div>
                <div className="location__type">{location.Type}</div>
            </div>
        );
    };

    return (
        <div className="locations">
            {locations.map((l) => <Location location={l} />)}
        </div>
    );

}