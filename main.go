package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/RuthCodina/APIRestGo/entities"
	"github.com/gorilla/mux"
)

var tasks = entities.AllTasks{
	{
		ID:      1,
		Name:    "Uno",
		Content: "Something",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Print(w, "welcome to the API")
}
func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}
func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // obtiene los params  y query params usados en la url de la petición.
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintf(w, "Error convirtiendo en int el param id")
	}

	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("content-type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}
func createTask(w http.ResponseWriter, r *http.Request) {
	newTask := entities.Task{}
	req, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Inserte la información correcta")
	}
	json.Unmarshal(req, &newTask)
	fmt.Println(string(req[:]))
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // obtiene los params  y query params usados en la url de la petición.
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintf(w, "Error convirtiendo en int el param id")
	}

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "La tarea con ID %v, ha sido eliminado satisfactoriamente", taskID)
		}
	}
}
func updateTask(w http.ResponseWriter, r *http.Request) {
	updatedTask := entities.Task{}
	params := mux.Vars(r) // obtiene los params  y query params usados en la url de la petición.
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Fprintf(w, "Error convirtiendo en int el param id")
	}
	req, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte la información correcta")
	}
	json.Unmarshal(req, &updatedTask)
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updatedTask.ID = taskID
			tasks = append(tasks, updatedTask)
			fmt.Fprintf(w, "La tarea con el ID %v ha sido actualizada", taskID)
		}
	}

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")       // por que el id cambia, entonces lo ponemos entre {}
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE") // por que el id cambia, entonces lo ponemos entre {}
	router.HandleFunc("/tasks/{id}", updateTask).Methods("PATCH")  // por que el id cambia, entonces lo ponemos entre {}
	router.HandleFunc("/tasks", createTask).Methods("POST")
	log.Fatalln(http.ListenAndServe(":3000", router))
}
