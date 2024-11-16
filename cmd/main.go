package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"

	"github.com/fletcharoo/snest"
)

type Config struct {
	Port string `snest:"PORT"`
}

//go:embed style.css
var styleCSS string

func main() {
	// Load service configs.
	var conf Config
	if err := snest.Load(&conf); err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	// Add style.css route.
	http.HandleFunc("/style.css", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Add("Content-Type", "text/css")
		fmt.Fprintf(w, styleCSS)
	})
	log.Println("Registered endpoint /style.css")

	// Start service.
	addr := ":" + conf.Port
	log.Println("Serving on", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
