package dtos

import "time"

type LocationDto struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Dimension string    `json:"dimension"`
	Residents []string  `json:"residents"`
	URL       string    `json:"url"`
	Created   time.Time `json:"created"`
}

type InfoDto struct {
	Count int    `json:"count"`
	Pages int    `json:"pages"`
	Next  string `json:"next"`
	Prev  string `json:"prev"`
}

type AllLocationsDto struct {
	Info    InfoDto       `json:"info"`
	Results []LocationDto `json:"results"`
}
