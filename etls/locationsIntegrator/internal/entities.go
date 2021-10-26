package entities

import "time"

type Entity struct {
	Id int
}

type Location struct {
	Entity
	Name      string    `bson:"Name"`
	Type      string    `bson:"Type"`
	Dimension string    `bson:"Dimension"`
	Residents []string  `bson:"Residents"`
	URL       string    `bson:"Url"`
	Created   time.Time `bson:"Created"`
}
