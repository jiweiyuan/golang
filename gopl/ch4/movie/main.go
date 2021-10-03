package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}




func main(){

	var movies = []Movie{
		{Title: "Casablanca", Year:1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergmann"}},
		{Title: "Cool Hand Lucky", Year:1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen","Jacqueline"}},
	}


	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)

	indentData, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("Json MarshalIndent failed: %s", err)
	}

	fmt.Printf("%s\n", indentData)

	var titles []struct{Title string}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON Unmarshaling failed %s\n", err)
	}
	fmt.Println(titles)

}
