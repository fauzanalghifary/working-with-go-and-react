package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// JSONResponse is a generic struct for sending JSON responses
type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)

	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)

	if err != nil {
		return err
	}

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {

	maxBytes := 1024 * 1024 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// decode the request body into the target destination
	dec := json.NewDecoder(r.Body)

	dec.DisallowUnknownFields()

	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil

}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	// if the status code is not provided, use http.StatusBadRequest
	statusCode := http.StatusBadRequest

	// if the status code is provided, use that instead
	if len(status) > 0 {
		statusCode = status[0]
	}

	// create a response payload to send to the client
	payload := JSONResponse{
		Error:   true,
		Message: err.Error(),
	}

	// send the response
	return app.writeJSON(w, statusCode, payload)
}
