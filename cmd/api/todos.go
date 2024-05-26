package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/manujelko/gotodo/internal/data"
)

func (app *application) getTodoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rawId := ps.ByName("id")
	id, err := strconv.ParseInt(rawId, 10, 64)
	if err != nil || id < 1 {
		errMsg := `"message": "Bad Id"`
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	todo, err := app.dao.Todos.Get(id)
	if err != nil {
		app.logger.Error(err.Error())
		errMsg := `"message": "Todo Not Found"`
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}
	js, err := json.Marshal(todo)
	if err != nil {
		app.logger.Error(err.Error())
		errMsg := `"message": "Internal Server Error"`
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func (app *application) getTodosHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		limit  int = 100
		offset int = 0
		err    error
	)

	rawLimit := r.URL.Query().Get("limit")
	rawOffset := r.URL.Query().Get("offset")

	if rawLimit != "" {
		limit, err = strconv.Atoi(rawLimit)
		if err != nil || limit < 1 {
			errMsg := `"message": "Bad Limit"`
			http.Error(w, errMsg, http.StatusBadRequest)
			return
		}
	}

	if rawOffset != "" {
		offset, err = strconv.Atoi(rawOffset)
		if err != nil || offset < 0 {
			errMsg := `"message": "Bad Offset"`
			http.Error(w, errMsg, http.StatusBadRequest)
			return
		}
	}

	todos, err := app.dao.Todos.GetMultiple(limit, offset)
	if err != nil {
		app.logger.Error(err.Error())
		errMsg := `"message": "Internal Server Error"`
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	data := map[string]any{"todos": todos}
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Error(err.Error())
		errMsg := `"message": "Internal Server Error"`
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(js)
}

func (app *application) createTodoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var input struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		errMsg := `"message": "Bad Request Body"`
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	todo := data.Todo{
		Title:   input.Title,
		Content: input.Content,
	}
	app.dao.Todos.Insert(&todo)
	js, err := json.Marshal(todo)
	if err != nil {
		app.logger.Error(err.Error())
		errMsg := `"message": "Internal Server Error"`
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func (app *application) completeTodoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rawId := ps.ByName("id")
	id, err := strconv.ParseInt(rawId, 10, 64)
	if err != nil || id < 1 {
		errMsg := `"message": "Bad Id"`
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	err = app.dao.Todos.Complete(id)
	if err != nil {
		app.logger.Error(err.Error())
		errMsg := `"message": "Internal Server Error"`
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	data := map[string]string{"message": "ok"}
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Error(err.Error())
		errMsg := `"message": "Internal Server Error"`
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func (app *application) deleteTodoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rawId := ps.ByName("id")
	id, err := strconv.ParseInt(rawId, 10, 64)
	if err != nil || id < 1 {
		http.Error(w, "Bad Id", http.StatusBadRequest)
		return
	}
	err = app.dao.Todos.Delete(id)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := map[string]string{"message": "ok"}
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
