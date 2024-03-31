package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var templates = template.Must(template.ParseGlob("web/templates/*"))

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveTemplate(w, r)
	})
	// New handler for search
	http.HandleFunc("/search", handleSearch)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/" {
		path = "/home"
	}
	tmpl, err := templates.Clone()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	_, err = tmpl.ParseFiles("web/templates" + path + ".html")
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"Title": strings.Title(strings.Trim(path, "/")),
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// handleSearch processes the search query and renders the search results page.
func handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	// Execute the Node.js search script
	cmd := exec.Command("node", "scripts/search.js", query)
	output, err := cmd.Output() // Use Output to capture standard output
	if err != nil {
		log.Printf("Failed to execute search script: %v", err)
		http.Error(w, "Search failed", http.StatusInternalServerError)
		return
	}

	// Assuming the script outputs JSON-formatted search results
	var results []map[string]interface{}
	if err := json.Unmarshal(output, &results); err != nil {
		log.Printf("Failed to parse search results: %v", err)
		http.Error(w, "Failed to parse search results", http.StatusInternalServerError)
		return
	}

	// Continue to render search results with your template
	tmpl, err := templates.Clone()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Assuming you have a "search_results.html" template for displaying the results
	_, err = tmpl.ParseFiles("web/templates/search_results.html")
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"Title":   "Search Results",
		"Query":   query,
		"Results": results, // Use the results from the Node.js script
	}

	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
