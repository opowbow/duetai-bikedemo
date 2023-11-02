package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	PageTitle  string
	MapsApiKey string
	Stations   []BikeStation
}

type BikeStation struct {
	Name               string
	Latitude           float64
	Longitude          float64
	BikesCount         int
	BikesVisualization string
}

func loadBikeStationData() []BikeStation {
	return []BikeStation{
		{
			Name:      "Google EURF",
			Latitude:  47.378867,
			Longitude: 8.5322847,
		},
	}
}

func main() {

	MapsApiKey := os.Getenv("MAPS_KEY")
	if MapsApiKey == "" {
		log.Fatal("Missing MAPS_API_KEY")
	}

	tmpl := template.Must(template.ParseFiles("template.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data := PageData{
			PageTitle:  "Pedal Power: London Bike Stations",
			MapsApiKey: MapsApiKey,
			Stations:   loadBikeStationData(),
		}

		tmpl.Execute(w, data)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8084"
		log.Printf("Running app on http://0.0.0.0:%s", port)
	}
	http.ListenAndServe(":"+port, nil)
}
