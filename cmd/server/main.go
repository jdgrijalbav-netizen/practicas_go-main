package main

import (
	"fmt"
	"log"
	"net/http"

	"practicas_go/internal/analyzer"
	"practicas_go/internal/api"
	"practicas_go/internal/handlers"
)

func main() {
	client := api.NewClient()
	analyzer := analyzer.NewAnalyzer(client)

	handler, err := handlers.NewHandler(analyzer)
	if err != nil {
		log.Fatalf("Failed to initialize handler: %v", err)
	}

	// Rutas principales
	http.HandleFunc("/", handler.ShowForm)
	http.HandleFunc("/analyze", handler.AnalyzeDomain)

	
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	fmt.Println("Server started at http://localhost:8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
