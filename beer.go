package main

type Beer struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Type	  string      `json:"type"`
	Kicked    bool   `json:"kicked"`
}

type Beers []Beer
