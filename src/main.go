package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Tool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	playwright := Tool{
		Name:        "Playwright",
		Description: "A web testing framework.",
	}
	tmpl := template.Must(template.ParseFiles("index.html"))

	// TODO - add input and submit which put these params in the query string
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tool := Tool{
			Name:        r.URL.Query().Get("name"),
			Description: r.URL.Query().Get("description"),
		}

		if tool.Name != "" {
			tmpl.Execute(w, tool)
		} else {
			tmpl.Execute(w, playwright)
		}
	})

	http.HandleFunc("/basic-rest", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(playwright)
	})

	// todo - basic-rest with auth

	// todo - chain of actions in REST
	http.HandleFunc("/action-chain-rest", func(w http.ResponseWriter, r *http.Request) {
		var tool Tool
		json.NewDecoder(r.Body).Decode(&tool)

		fmt.Fprintf(w, "%s: %s", tool.Name, tool.Description)
	})

	http.ListenAndServe(":8080", nil)
}
