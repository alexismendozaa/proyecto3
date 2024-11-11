package main

import (
	"encoding/json"
	"net/http"
)

type Tarea struct {
	ID   string `json:"id"`
	Desc string `json:"desc"`
}

var tareas []Tarea

func main() {
	// Agregar algunas tareas de ejemplo
	tareas = append(tareas, Tarea{ID: "1", Desc: "Estudiar Docker"})
	tareas = append(tareas, Tarea{ID: "2", Desc: "Leer documentación de Go"})
	tareas = append(tareas, Tarea{ID: "3", Desc: "Practicar Go y Docker"})

	// Manejar la ruta /tareas
	http.HandleFunc("/tareas", getTareas)

	// Escuchar en el puerto 8080
	http.ListenAndServe(":8080", nil)
}

func getTareas(w http.ResponseWriter, r *http.Request) {
	// Codificar las tareas en formato JSON y enviarlas como respuesta
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tareas)
}
