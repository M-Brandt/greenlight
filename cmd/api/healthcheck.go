package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	enve := envelope{
		"status": "available",
		"sytem_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, enve, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
