package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	// healthcheck
	router.GET("/v1/healthcheck", app.healthcheckHandler)

	// todos
	router.GET("/v1/todos/:id", app.getTodoHandler)
	router.GET("/v1/todos", app.getTodosHandler)
	router.POST("/v1/todos", app.createTodoHandler)
	router.PATCH("/v1/todos/:id", app.completeTodoHandler)
	router.DELETE("/v1/todos/:id", app.deleteTodoHandler)

	return router
}
