package server

import (
	"net/http"
	"net/url"

	"github.com/freakkid/service-agenda/service/entities"
	"github.com/unrolled/render"
)

// get user key by username and password, need key
func userLoginHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm() // parsing the parameters
		sessionID, status, responseJSON := entities.AgendaService.LoginAndGetSessionID(req.FormValue("username"), req.FormValue("password"))
		if sessionID != "" || status == http.StatusOK {
			http.SetCookie(w, &http.Cookie{Name: req.FormValue("username"), Value: url.QueryEscape(sessionID)})
		}
		formatter.JSON(w, status, responseJSON)
	}
}

// get user key by username and password, need key
func userLogoutHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()                                    // parsing the parameters
		cookie, _ := req.Cookie(req.FormValue("username")) // get cookie
		if cookie == nil {
			formatter.JSON(w, http.StatusUnauthorized, entities.SingleMessageResponse{Message: "please sign in to Agenda"})
		} else {
			status, responseJSON := entities.AgendaService.LogoutAndDeleteSessionID(cookie.Value)
			formatter.JSON(w, status, responseJSON)
		}
	}
}

// list limit users or get a user by id, need key
func usersInfoHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()                                    // parsing the parameters
		cookie, _ := req.Cookie(req.FormValue("username")) // get cookie
		if cookie == nil {
			formatter.JSON(w, http.StatusUnauthorized, entities.SingleMessageResponse{Message: "please sign in to Agenda"})
		} else {
			if req.FormValue("id") != "" {
				status, responseJSON := entities.AgendaService.GetUserInfoByID(cookie.Value, req.FormValue("id"))
				formatter.JSON(w, status, responseJSON)
			} else {
				status, responseJSON := entities.AgendaService.ListUsersByLimit(cookie.Value, req.FormValue("limit"), req.FormValue("offset"))
				formatter.JSON(w, status, responseJSON)
			}
		}
	}
}

// create a new user by username, password, email, password, no need key
func createUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm() // parsing the parameters
		status, responseJSON := entities.AgendaService.CreateUser(
			req.FormValue("username"), req.FormValue("password"), req.FormValue("email"), req.FormValue("phone"))
		formatter.JSON(w, status, responseJSON)
	}
}

// delete a user by password, need key
func deleteUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()                                    // parsing the parameters
		cookie, _ := req.Cookie(req.FormValue("username")) // get cookie
		if cookie == nil {
			formatter.JSON(w, http.StatusUnauthorized, entities.SingleMessageResponse{Message: "please sign in to Agenda"})
		}
		status, responseJSON := entities.AgendaService.DeleteUserByPassword(cookie.Value, req.FormValue("password"))
		formatter.JSON(w, status, responseJSON)
	}
}
