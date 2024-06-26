package main

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blues", Artist: "J C", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "G M", Price: 17.99},
}
