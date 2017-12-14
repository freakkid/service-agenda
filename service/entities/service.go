package entities

import (
	"net/http"
	"strconv"

	"github.com/freakkid/service-agenda/service/tools"
)

// AgendaAtomicService -- a struct to operate service function
type AgendaAtomicService struct{}

// AgendaService -- an instance
var AgendaService = AgendaAtomicService{}

//
// ─── TO BE JSON RESPONSE ───────────────────────────────────────────────────────────
//

// UserKeyResponse -- GetUserByKeyAndID
type UserKeyResponse struct {
	Message  string
	ID       int
	UserName string
	Email    string
	Phone    string
}

// SingleMessageResponse -- DeleteUserByKeyAndPassword/ChangeUserPassword/GetUserKey
type SingleMessageResponse struct {
	Message string
}

// used in UsersInfoResponse
type singleUserInfo struct {
	ID       int
	UserName string
	Email    string
	Phone    string
}

// UsersInfoResponse -- ListUsersByKeyAndLimit
type UsersInfoResponse struct {
	Message            string
	SingleUserInfoList []singleUserInfo
}

// CreateUserResponse -- CreateUser
type CreateUserResponse struct {
	Message  string
	ID       int
	UserName string
	Email    string
	Phone    string
}

// ────────────────────────────────────────────────────────────────────────────────

//
// ─── PROVIDE SERICES AND RETURN STATUS CODE AND JSON RESPONSE STRUCT ────────────
//

// CreateUser -- check if input is empty and username is duplicate
func (*AgendaAtomicService) CreateUser(
	username string, password string, email string, phone string) (int, CreateUserResponse) {
	// ---- check input ----
	if username == "" || password == "" || email == "" || phone == "" {
		return http.StatusBadRequest, CreateUserResponse{Message: "empty input", ID: -1}
	}
	password = tools.MD5Encryption(password)
	dao := agendaDao{xormEngine}
	// ---- check username ----
	has, err := dao.ifUserExistByConditions(&User{UserName: username})
	if err != nil { // server error
		return http.StatusInternalServerError, CreateUserResponse{Message: "server error", ID: -1}
	}
	if has { // username exist -- duplicate username
		return http.StatusBadRequest, CreateUserResponse{Message: "duplicate username", ID: -1}
	}
	// ---- create user ----
	result, user := dao.createUser(&User{SessionID: tools.GetKey(), UserName: username, Password: password, Email: email, Phone: phone})
	if result && user != nil { // create user successfully
		return http.StatusCreated, CreateUserResponse{"create user " + username + " successfully",
			user.ID, user.UserName, user.Email, user.Phone}
	}
	return http.StatusBadRequest, CreateUserResponse{Message: "maybe username is duplicate", ID: -1}
}

// LoginAndGetSessionID --- check if user exists and generate key
// if user no exists or occur error, return empty string and error
// if get key success, return key and empty error
func (*AgendaAtomicService) LoginAndGetSessionID(username string, password string) (string, int, SingleMessageResponse) {
	// ---- check GET data ----
	if username == "" || password == "" { // check if empty username and password
		return "", http.StatusBadRequest, SingleMessageResponse{"empty username and password"}
	}
	password = tools.MD5Encryption(password)
	dao := agendaDao{xormEngine}
	has, err := dao.ifUserExistByConditions(&User{UserName: username, Password: password})
	if err != nil { // server error
		return "", http.StatusInternalServerError, SingleMessageResponse{"server error"}
	}
	if !has { // user not exist
		return "", http.StatusUnauthorized, SingleMessageResponse{"incorrect username or password"}
	}
	// ---- get new sessionID ----
	var sessionID = tools.GenenrateSessionID() // generate new sessionID
	for {                                      // make sure sessionID unique
		if has, err = dao.ifUserExistByConditions(&User{SessionID: sessionID}); err == nil && !has {
			break
		}
		sessionID = tools.GenenrateSessionID()
	}
	affected, _ := dao.updateUser(&User{SessionID: sessionID}, &User{UserName: username, Password: password})
	if affected == 0 { // user not exist
		return "", http.StatusUnauthorized, SingleMessageResponse{"incorrect username or password"}
	}
	return sessionID, http.StatusOK, SingleMessageResponse{"login successfully"}
}

// GetUserInfoByID --- convert string id to int id, if occur error return empty user and error
// check if key is valid and id exsits and belong to the same user
// if valid key and exist id, return User struct
func (*AgendaAtomicService) GetUserInfoByID(key string, stringID string) (int, UserKeyResponse) {
	var (
		id   int
		err  error
		has  bool
		user *User
	)
	dao := agendaDao{xormEngine}
	// ---- check key ----
	has, err = dao.ifUserExistByConditions(&User{SessionID: key})
	if err != nil { // server error
		return http.StatusInternalServerError, UserKeyResponse{Message: "server error", ID: -1}
	}
	if !has { // invalid key
		return http.StatusUnauthorized, UserKeyResponse{Message: "invalid key", ID: -1}
	}
	// ---- check id ----
	if stringID == "" { // empty id
		return http.StatusBadRequest, UserKeyResponse{Message: "empty id", ID: -1}
	}
	id, err = strconv.Atoi(stringID)
	if err != nil || id <= 0 { // invalid id
		return http.StatusBadRequest, UserKeyResponse{Message: "invalid id", ID: -1}
	}
	// ---- find user by id ----
	has, user = dao.findUserByConditions(&User{ID: id})
	if has && user != nil { // user not exist
		return http.StatusOK,
			UserKeyResponse{"get user info successfully", user.ID, user.UserName, user.Email, user.Phone}
	}
	return http.StatusNotFound,
		UserKeyResponse{Message: "the user with id " + stringID + " not exists", ID: id}
}

// DeleteUserByPassword --- check key if valid
// check if password correct
func (*AgendaAtomicService) DeleteUserByPassword(key string, password string) (int, SingleMessageResponse) {
	var (
		err      error
		has      bool
		affected int64
	)
	dao := agendaDao{xormEngine}
	// ---- check key ----
	has, err = dao.ifUserExistByConditions(&User{SessionID: key})
	if err != nil { // server error
		return http.StatusInternalServerError, SingleMessageResponse{Message: "server error"}
	}
	if !has { // invalid key
		return http.StatusUnauthorized, SingleMessageResponse{Message: "invalid key"}
	}
	// ---- check password ----
	if password == "" { // empty input
		return http.StatusBadRequest, SingleMessageResponse{"empty password"}
	}
	affected, err = dao.deleteUserByKeyAndPassword(key, tools.MD5Encryption(password))
	if err != nil { // server error
		return http.StatusInternalServerError, SingleMessageResponse{Message: "server error"}
	}
	if affected == 0 { // delete user fail
		return http.StatusUnauthorized, SingleMessageResponse{Message: "incorrect paassword"}
	}
	// delete successfully
	return http.StatusNoContent, SingleMessageResponse{}
}

// ListUsersByLimit --- check key is valid or not
// if limit is invalid, default set to 10
func (*AgendaAtomicService) ListUsersByLimit(key string, stringLimit string, stringOffset string) (int, UsersInfoResponse) {
	var (
		limit int
		has   bool
		err   error
		users []User
	)
	dao := agendaDao{xormEngine}
	// ---- check key ----
	has, err = dao.ifUserExistByConditions(&User{SessionID: key})
	if err != nil { // server error
		return http.StatusInternalServerError, UsersInfoResponse{"server error", []singleUserInfo{}}
	}
	if !has { // if key not exist -- invalid key
		return http.StatusUnauthorized, UsersInfoResponse{"invalid key", []singleUserInfo{}}
	}
	// ---- check limit ----
	if stringLimit == "" {
		limit = 5
	} else {
		limit, err = strconv.Atoi(stringLimit)
		if err != nil || limit <= 0 { // invalid limit
			return http.StatusBadRequest, UsersInfoResponse{"invalid limit", []singleUserInfo{}}
		}
	}
	// ---- get limit users ----
	users, err = dao.getLimitUsers(limit)
	if err != nil { // server error
		return http.StatusInternalServerError, UsersInfoResponse{"server error", []singleUserInfo{}}
	}
	singleUserInfoList := make([]singleUserInfo, 0, 0)
	for _, userInfo := range users {
		singleUserInfoList = append(singleUserInfoList,
			singleUserInfo{userInfo.ID, userInfo.UserName, userInfo.Email, userInfo.Phone})
	}
	return http.StatusOK, UsersInfoResponse{"get userlist successfully", singleUserInfoList}
}

// ChangeUserPassword -- check if key is valid
// check if password is correct
// check if new password valid and match confirmation
func ChangeUserPassword(key string, password string, newPassword string, confirmation string) (int, SingleMessageResponse) {
	var (
		has      bool
		err      error
		affected int64
	)
	dao := agendaDao{xormEngine}
	// ---- check key ----
	has, err = dao.ifUserExistByConditions(&User{SessionID: key})
	if err != nil { // server error
		return http.StatusInternalServerError, SingleMessageResponse{"server error"}
	}
	if !has { // if key not exist -- invalid key
		return http.StatusUnauthorized, SingleMessageResponse{"invalid key"}
	}
	// ---- check old password ----
	has, err = dao.ifUserExistByConditions(&User{SessionID: key, Password: password})
	if err != nil { // server error
		return http.StatusInternalServerError, SingleMessageResponse{"server error"}
	}
	if !has { // password incorrect
		return http.StatusUnauthorized, SingleMessageResponse{"incorrect password"}
	}
	// ---- check new password ----
	if newPassword == "" {
		return http.StatusBadRequest, SingleMessageResponse{"new password is empty"}
	}
	if newPassword != confirmation {
		return http.StatusBadRequest, SingleMessageResponse{"new password and confirmation do not match"}
	}
	// ---- update new password ----
	affected, _ = dao.updateUser(&User{Password: newPassword}, &User{SessionID: key, Password: password})
	if affected == 0 { // user not exist
		return http.StatusUnauthorized, SingleMessageResponse{"incorrect password"}
	}
	return http.StatusOK, SingleMessageResponse{"update password successfully"}
}

// LogoutAndDeleteSessionID -- logout and update sessionid
func LogoutAndDeleteSessionID(sessionID string) (int, SingleMessageResponse) {
	// ---- check sessionID ----
	if sessionID == "" { // check if empty sessionID
		return http.StatusUnauthorized, SingleMessageResponse{"log out fail"}
	}
	dao := agendaDao{xormEngine}
	has, err := dao.ifUserExistByConditions(&User{SessionID: sessionID})
	if err != nil { // server error
		return http.StatusInternalServerError, SingleMessageResponse{"server error"}
	}
	if !has { // user not exist
		return http.StatusUnauthorized, SingleMessageResponse{"log out faile"}
	}
	// ---- get new sessionID ----
	var newSessionID = tools.GenenrateSessionID() // generate new sessionID
	for {                                         // make sure new sessionID unique
		if has, err = dao.ifUserExistByConditions(&User{SessionID: newSessionID}); err == nil && !has {
			break
		}
		newSessionID = tools.GenenrateSessionID()
	}
	affected, _ := dao.updateUser(&User{SessionID: newSessionID}, &User{SessionID: sessionID}) // replace old sessionID
	if affected == 0 {                                                                         // user not exist
		return http.StatusUnauthorized, SingleMessageResponse{"log out fail"}
	}
	return http.StatusOK, SingleMessageResponse{"log out successfully"}
}
