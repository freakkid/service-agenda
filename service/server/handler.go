package server

import (
	"net/http"

	"github.com/freakkid/service-agenda/service/entities"
	"github.com/unrolled/render"
)

// get user key by username and password, need key
func userGetKeyHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm() // parsing the parameters
		status, responseJSON := entities.AgendaService.GetUserKey(req.FormValue("username"), req.FormValue("password"))
		formatter.JSON(w, status, responseJSON)
	}
}

// list limit users or get a user by id, need key
func usersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm() // parsing the parameters
		if req.FormValue("id") != "" {
			status, responseJSON := entities.AgendaService.GetUserByKeyAndID(req.FormValue("key"), req.FormValue("id"))
			formatter.JSON(w, status, responseJSON)
		} else {
			status, responseJSON := entities.AgendaService.ListUsersByKeyAndLimit(req.FormValue("key"), req.FormValue("limit"))
			formatter.JSON(w, status, responseJSON)
		}
	}
}

// create a new user by username, password, email, password, no need key
func createUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm() // parsing the parameters
		status, responseJSON := entities.AgendaService.CreateUser(req.FormValue("username"), req.FormValue("password"), req.FormValue("email"), req.FormValue("phone"))
		formatter.JSON(w, status, responseJSON)
	}
}

// delete a user by password, need key
func deleteUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm() // parsing the parameters
		status, responseJSON := entities.AgendaService.DeleteUserByKeyAndPassword(req.FormValue("key"), req.FormValue("password"))
		formatter.JSON(w, status, responseJSON)
	}
}
