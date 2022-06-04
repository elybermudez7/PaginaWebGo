package main

import (
	"html/template"
	"log"
	"net/http"
)

var plantilla = template.Must(template.ParseGlob("Plantilla/panel/*.html"))

func index(w http.ResponseWriter, req *http.Request) {
	plantilla.ExecuteTemplate(w, "index.html", nil)
}

func princing(w http.ResponseWriter, req *http.Request) {
	plantilla.ExecuteTemplate(w, "princing.html", nil)
}

func about(w http.ResponseWriter, req *http.Request) {
	plantilla.ExecuteTemplate(w, "about.html", nil)
}

func testimonial(w http.ResponseWriter, req *http.Request) {
	plantilla.ExecuteTemplate(w, "testimonial.html", nil)
}

func why(w http.ResponseWriter, req *http.Request) {
	plantilla.ExecuteTemplate(w, "why.html", nil)
}

func main() {

	css := http.FileServer(http.Dir("elementos"))

	http.Handle("/elementos/", http.StripPrefix("/elementos", css))

	log.Println("Iniciando el servidor...")
	log.Println("Cargando las paginas web...")
	log.Println("Servidor iniciado en http://localhost:8080")

	http.HandleFunc("/", index)
	http.HandleFunc("/about", info)

	http.ListenAndServe(":8080", nil)
}
