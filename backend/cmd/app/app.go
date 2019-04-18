package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	todo struct {
		ID      int    `json:"id"`
		Title   string `json:"title"`
		Checked bool   `json:"checked"`
	}
)

var (
	todos = map[int]*todo{}
	seq   = 1
)

func createTodo(c echo.Context) error {
	t := &todo{
		ID:      seq,
		Checked: false,
	}
	if err := c.Bind(t); err != nil {
		return err
	}
	todos[t.ID] = t
	seq++
	return c.JSON(http.StatusCreated, t)
}

func getTodo(c echo.Context) error {
	return c.JSON(http.StatusOK, todos)
}

func updateTodo(c echo.Context) error {
	t := new(todo)
	if err := c.Bind(t); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	todos[id].Checked = t.Checked
	return c.JSON(http.StatusOK, todos[id])
}

func deleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(todos, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	// Routes
	e.POST("/todos", createTodo)
	e.GET("/todos", getTodo)
	e.PUT("/todos/:id", updateTodo)
	e.DELETE("/todos/:id", deleteTodo)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
