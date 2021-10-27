export interface OriginType {
    Name: string;
    URL: string;
}

export interface LocationType {
    Name: string;
    URL: string;
}

export interface CharacterType {
    Id: number;
    Name: string;
    Status: string;
    Species: string;
    Type: string;
    Gender: string;
    Origin: OriginType;
    Location: LocationType;
    Image: string;
    Episode: string[];
    URL: string;
    Created: Date;
}
