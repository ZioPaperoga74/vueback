package main

import (
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		UserName string `json:"email"`
		Password string `json:"password"`
	}

	var creds credentials
	var payload jsonResponse

	err := app.readJSON(w, r, &creds)
	if err != nil {
		app.errorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid json supplied or json missed entirely"

		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	// err := json.NewDecoder(r.Body).Decode(&creds)
	// if err != nil {
	// 	// send back msg
	// 	app.errorLog.Println("Invalid json")
	// 	payload.Error = true
	// 	payload.Message = "Invalid json"

	// 	out, err := json.MarshalIndent(payload, "", "\t")
	// 	if err != nil {
	// 		app.errorLog.Println(err)
	// 	}
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write(out)
	// 	return

	// }

	// todo authenticate
	app.infoLog.Println(creds.UserName, creds.Password)

	// send back a response

	payload.Error = false
	payload.Message = "Login successful"

	// out, err := json.MarshalIndent(payload, "", "\t")
	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}
