// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package wineo_api_service

type Wine struct {
	Winery     string  `json:"winery" gorethink:"id[0]"`
	Varietal   string  `json:"varietal" gorethink:"id[1]"`
	Vintage    int     `json:"vintage" gorethink:"id[2]"`
	BottleSize int     `json:"bottleSize" gorethink:"id[3]"`
	Region     *string `json:"region"`
	Peak       *int    `json:"peak"`
	Hold       *int    `json:"hold"`
	Drink      *int    `json:"drink"`
	Taste      *int    `json:"taste"`
	Price      *int    `json:"price"`
}