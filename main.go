package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

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
func createTask(w http.ResponseWriter, r *http.Request) {
	newTask := entities.Task{}
	req, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Inserte la información correcta")
	}
	json.Unmarshal(req, &newTask)
	fmt.Println(newTask)
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	log.Fatalln(http.ListenAndServe(":3000", router))
}
