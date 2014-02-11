package main

import (
	"net/http"
)

func main() {
	// http.Handle("/events", handler)
	// http.Handle("/diagnostic", handler) // diagnostic handler
	http.Handle("/dashboard/", http.StripPrefix("/assets", http.FileServer(http.Dir("public"))))

	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("public"))))
	http.Handle("/", http.RedirectHandler("/dashboard", http.StatusFound))
	http.ListenAndServe(":8080", nil)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
