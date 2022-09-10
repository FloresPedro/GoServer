package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Curso struct {
	Title        string `json:"title"`
	NumeroVideos int    `json:"numero_videos"`
}

type Cursos []Curso

func main() {
	//rutas
	http.HandleFunc("/holamundo", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hola pirus, entrando a un url mas chida")
	})

	//servir archivos estaticos
	http.HandleFunc("/archivo/estatico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	//Respoder formato JSON
	http.HandleFunc("/formatoJson", handlerJSON)

	//ejemplo que devuelve una estructura de cursos en formato JSON
	http.HandleFunc("/formatoJson/data-curso", handlerFunData)

	//ejemplo para un request
	http.HandleFunc("/", handler)

	//donde escucha el server
	http.ListenAndServe(":8000", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hay una nueva peticion")
	io.WriteString(w, "Hola Mundo Web")
}

// funcio para responder formato JSON
func handlerJSON(w http.ResponseWriter, r *http.Request) {
	curso := Curso{"Curso de Go", 30}
	json.NewEncoder(w).Encode(curso)
}

// response para la data Curso
func handlerFunData(w http.ResponseWriter, r *http.Request) {
	cursos := Cursos{Curso{"Curso de Go", 30}, Curso{"Curso de Java", 40}}
	json.NewEncoder(w).Encode(cursos)
}
