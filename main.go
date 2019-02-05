package main

import (
    "encoding/json"
    "log"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"fmt"
)
type Tasks struct {
    ID 			string  	`json:"id,omitempty"`
    TASKNAME 	string 		`json:"task,omitempty"`
}

var task []Tasks


func getAllTask(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(task)
}
func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _,item := range task {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Tasks{})
}

func createTask(w http.ResponseWriter, r *http.Request) {
    //TODO
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range task {
        if item.ID == params["id"] {
            task = append(task[:index], task[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(task)
    }
}

func homepage(w http.ResponseWriter, r *http.Request)  {
	json.NewEncoder(w).Encode("API is live..")
	
}


func main() {

	task = append(task, Tasks{ID: "1", TASKNAME: "clean room"})
	task = append(task, Tasks{ID: "2", TASKNAME: "watch movie"})
	task = append(task, Tasks{ID: "3", TASKNAME : "eat oreo"})

	router := mux.NewRouter()
	router.HandleFunc("/", homepage).Methods("GET")
	router.HandleFunc("/task", getAllTask).Methods("GET")
	router.HandleFunc("/task/{id}", getTask).Methods("GET")
	router.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/task/{id}", createTask).Methods("POST")

    headersOk := handlers.AllowedHeaders([]string{"Authorization"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

    fmt.Println("Server running at PORT 4200....")
    log.Fatal(http.ListenAndServe(":4200", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

