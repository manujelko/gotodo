package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	data := map[string]any{
		"status": "available",
		"system_information": map[string]string{
			"version":     version,
			"environment": app.cfg.env,
		},
	}
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
