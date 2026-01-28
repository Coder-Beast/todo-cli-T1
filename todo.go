package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var todos =  []Todo{
	{"01","Buy Milk", false},
	{"02","Go Walking", false},
	{"03","Eat", false},
}

func getTodos(w http.ResponseWriter, r *http.Request){


    rows, err := db.Query("SELECT id, item, completed FROM todos")
    if err != nil{
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }

    defer rows.Close()

    var todoList []Todo

    for rows.Next(){
        var t Todo
        err := rows.Scan(&t.Id,&t.Item,&t.Completed)
        if err != nil {
            log.Println(err)
            return
        }

        todoList = append(todoList, t)
    }
    if todoList == nil{
        todoList = []Todo{}
    }
    respondJSON(w, http.StatusOK, todoList)
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
    }
   

    query := "INSERT INTO todos (id, item, completed) VALUES (?, ?, ?)"

    _,err = db.Exec(query,newTodo.Id,newTodo.Item,newTodo.Completed)
    if err != nil {
       respondError(w, http.StatusInternalServerError, "Failed to save to database")
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newTodo)

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

    result, err := db.Exec("DELETE FROM todos WHERE id = ?", target.Id)
    if err!=nil{
        http.Error(w,"Failed to delete", 500)
        return
    }
    rowsaffected, _:= result.RowsAffected()

    if rowsaffected == 0{
        respondError(w,404,"Target not found")
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode("Deleted Successfully")
    /*for index,item := range todos{
        if item.Id == target.Id{
            todos = append(todos[:index], todos[index+1:]...)
            
            return
        }
    }*/

}