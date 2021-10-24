package entities

import "time"

type Location struct {
	Id        int       `bson:"Id"`
	Name      string    `bson:"Name"`
	Type      string    `bson:"Type"`
	Dimension string    `bson:"Dimension"`
	Residents []string  `bson:"Residents"`
	URL       string    `bson:"Url"`
	Created   time.Time `bson:"Created"`
}
