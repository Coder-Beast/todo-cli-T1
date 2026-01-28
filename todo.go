package main

import (
	"net/http"
	"encoding/json"
    "fmt"
)

var todos =  []Todo{
	{"01","Buy Milk", false},
	{"02","Go Walking", false},
	{"03","Eat", false},
}

func getTodos(w http.ResponseWriter, r *http.Request){
	respondJSON(w,http.StatusOK,todos)
    
}
func addTodo(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        respondError(w,http.StatusMethodNotAllowed,"Method not allowed")
        return
    }

    var newTodo Todo

    err := json.NewDecoder(r.Body).Decode(&newTodo)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }else{
        fmt.Println("Decoded")
    }

    todos = append(todos, newTodo)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newTodo)
    fmt.Println("Encded")

}



func deleteTodo(w http.ResponseWriter,r *http.Request){
    if r.Method != http.MethodDelete{
        http.Error(w, "Method not allowed Use Delete", http.StatusMethodNotAllowed)
        return
    }

    var target struct {
        Id string `json:"id"`
    }
    err := json.NewDecoder(r.Body).Decode(&target)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    for index,item := range todos{
        if item.Id == target.Id{
            todos = append(todos[:index], todos[index+1:]...)
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode("Deleted Successfully")
            return
        }
    }

    http.Error(w, "ID not found", http.StatusNotFound)
}