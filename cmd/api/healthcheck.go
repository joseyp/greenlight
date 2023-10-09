package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, req *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, req, err)
	}
}
