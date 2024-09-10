package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	data := struct {
		Title string
		Items []string
	}{
		Title: "Welcome bro",
		Items: []string{},
	}

	indexPageHandler := func(resWriter http.ResponseWriter, request *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))

		templ.Execute(resWriter, data)
	}

	changeMessageHandler := func(resWriter http.ResponseWriter, request *http.Request) {
		message := request.PostFormValue("message")

		// Prepare the HTML fragment to be swapped in
		response := fmt.Sprintf("<p>Message received: %s</p>", message)

		// Write the HTML fragment directly in the response
		resWriter.Header().Set("Content-Type", "text/html")
		resWriter.Write([]byte(response))
	}

	addItemHandler := func(resWriter http.ResponseWriter, request *http.Request) {
		newItem := request.PostFormValue("item")
		templ := template.Must(template.ParseFiles("index.html"))

		data.Items = append(data.Items, newItem)

		// optional, to set content type
		resWriter.Header().Set("Content-Type", "text/html")
		// this will evaluate the `item-list-element` subtemplate / block and return the HTML response
		templ.ExecuteTemplate(resWriter, "item-list-element", newItem)
	}

	http.HandleFunc("/", indexPageHandler)

	http.HandleFunc("/change-message", changeMessageHandler)
	http.HandleFunc("/add-item", addItemHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
