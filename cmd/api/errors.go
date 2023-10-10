package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(req *http.Request, err error) {
	app.logger.Print(err)
}

func (app *application) errorResponse(w http.ResponseWriter, req *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(req, err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, req *http.Request, err error) {
	app.logError(req, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, req, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, req *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, req, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, req *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", req.Method)
	app.errorResponse(w, req, http.StatusMethodNotAllowed, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, req *http.Request, err error) {
	app.errorResponse(w, req, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, req *http.Request, errors map[string]string) {
	app.errorResponse(w, req, http.StatusUnprocessableEntity, errors)
}
