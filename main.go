package main
import(
	"net/http"
	"fmt"
)
func main(){

	initDB()
	http.HandleFunc("/todos",getTodos)
	http.HandleFunc("/add-todo", addTodo)
	http.HandleFunc("/delete-todo", deleteTodo)
	fmt.Println("Server starting on port 9090...")

	err := http.ListenAndServe(":9090",nil)
	if err != nil {
        fmt.Println("Error starting server:", err)
    }
}