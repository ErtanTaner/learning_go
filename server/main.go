package main

import (
	"encoding/json"
	"example.com/net-test/docs"
	"example.com/net-test/types"
	"example.com/net-test/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"

	_ "example.com/net-test/docs"
	"github.com/swaggo/http-swagger/v2"
)

var todos []types.Todo = make([]types.Todo, 0)

// @Summary Create todo
// @Description Create a new todo
// @Tags todos
// @Accept json
// @Param todo body types.Todo true "Todo"
// @Success 201 {string} string
// @Router /todo [post]
func createTodo(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var todo types.Todo
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&todo)
	todos = append(todos, todo)
	encoder := json.NewEncoder(w)
	encoder.Encode(todo)
}

// @Summary Get todos
// @Description Get all todos or a single todo by ID (use id query param)
// @Tags todos
// @Produce json
// @Param id query string false "Todo ID (returns single todo if provided)"
// @Success 200 {array} types.Todo
// @Router /todo [get]
func getTodos(w http.ResponseWriter, id string) {
	if id != "" {
		idx := utils.FindIdx(todos, id)
		encoder := json.NewEncoder(w)

		encoder.Encode(todos[idx])
	} else {
		encoder := json.NewEncoder(w)

		encoder.Encode(todos)
	}
}

// @Summary Delete todo
// @Description Delete a todo by ID
// @Tags todos
// @Param id query string true "Todo ID"
// @Success 200 {string} string
// @Router /todo [delete]
func deleteTodo(w http.ResponseWriter, id string) {
	idx := utils.FindIdx(todos, id)
	todos = slices.Delete(todos, idx, idx+1)

	io.WriteString(w, "Todo deleted successfully!")
}

// @Summary Update todo
// @Description Update a todo by ID
// @Tags todos
// @Accept json
// @Param todo body types.Todo true "Todo"
// @Success 200 {string} string
// @Router /todo [put]
func updateTodo(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var reqTodo types.Todo
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&reqTodo)

	for i, v := range todos {
		if v.ID == reqTodo.ID {
			todo := &todos[i]
			todo.Name = reqTodo.Name
			todo.Done = reqTodo.Done
		}
	}

	io.WriteString(w, "Todo updated successfully!")
}

func todoHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if req.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")
		createTodo(w, req)
	} else if req.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		getTodos(w, id)
	} else if req.Method == "DELETE" {
		deleteTodo(w, id)
	} else if req.Method == "PUT" {
		updateTodo(w, req)
	}
}

// @Title Simple Todo API
// @Description Todo management API
func main() {
	http.HandleFunc("/todo", todoHandler)

	docs.SwaggerInfo.Title = "Simple HTTP server API"
	docs.SwaggerInfo.Description = "This is a sample server net/http server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	http.HandleFunc("/swagger/", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	fmt.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
