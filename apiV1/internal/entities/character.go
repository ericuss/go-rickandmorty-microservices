package entities

type Character struct {
	Id      int    `bson:"Id"`
	Name    string `bson:"Name"`
	Status  string `bson:"Status"`
	Species string `bson:"Species"`
	Type    string `bson:"Type"`
	Gender  string `bson:"Gender"`
	Origin  struct {
		Name string `bson:"Name"`
		URL  string `bson:"Url"`
	} `bson:"Origin"`
	Location struct {
		Name string `bson:"Name"`
		URL  string `bson:"Url"`
	} `bson:"Location"`
	Image   string   `bson:"Image"`
	Episode []string `bson:"Episode"`
	URL     string   `bson:"Url"`
	Created string   `bson:"Created"`
}
