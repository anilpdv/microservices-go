package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"microservicego/mvc/services"
	"microservicego/mvc/utils"
)

// GetUser : func (w http.ResponseWriter, r *http.Request)
func GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		userError := new(utils.ApplicationError)
		userError.ApplicationErrorCreator("user_id should be number", http.StatusBadRequest, "bad_request")
		userErrorMarshal, _ := json.Marshal(userError)
		w.WriteHeader(userError.StatusCode)
		w.Write(userErrorMarshal)
		return
	}

	user, appErr := services.UserService.GetUser(id)
	appErrMarshal, _ := json.Marshal(appErr)
	if appErr != nil {
		w.WriteHeader(appErr.StatusCode)
		w.Write(appErrMarshal)
		return
	}

	usermarshal, _ := json.Marshal(user)
	w.Write(usermarshal)
}
