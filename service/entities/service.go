package entities

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/freakkid/service-agenda/service/tools"
)

// AgendaAtomicService -- a struct to operate service function
type AgendaAtomicService struct{}

// AgendaService -- an instance
var AgendaService = AgendaAtomicService{}

// GetUserKeyResponse -- GetUserKey
type GetUserKeyResponse struct {
	Key     string
	Message string
}

// UserKeyResponse -- GetUserByKeyAndID
type UserKeyResponse struct {
	Message  string
	ID       int
	UserName string
	Email    string
	Phone    string
}

// DeleteUserResponse -- DeleteUserByKeyAndPassword
type DeleteUserResponse struct {
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
		fmt.Println(err)
		return http.StatusInternalServerError, CreateUserResponse{Message: "server error", ID: -1}
	}
	if has { // username exist -- duplicate username
		return http.StatusBadRequest, CreateUserResponse{Message: "duplicate username", ID: -1}
	}
	// ---- create user ----
	result, user := dao.createUser(&User{Key: tools.GetKey(), UserName: username, Password: password, Email: email, Phone: phone})
	if result && user != nil { // create user successfully
		return http.StatusCreated, CreateUserResponse{"create user " + username + " successfully",
			user.ID, user.UserName, user.Email, user.Phone}
	}
	return http.StatusBadRequest, CreateUserResponse{Message: "maybe username is duplicate", ID: -1}
}

// GetUserKey --- check if user exists and generate key
// if user no exists or occur error, return empty string and error
// if get key success, return key and empty error
func (*AgendaAtomicService) GetUserKey(username string, password string) (int, GetUserKeyResponse) {
	// ---- check GET data ----
	if username == "" || password == "" { // check if empty username and password
		return http.StatusBadRequest, GetUserKeyResponse{"", "empty username and password"}
	}
	password = tools.MD5Encryption(password)
	dao := agendaDao{xormEngine}
	has, err := dao.ifUserExistByConditions(&User{UserName: username, Password: password})
	if err != nil { // server error
		return http.StatusInternalServerError, GetUserKeyResponse{"", err.Error()}
	}
	if !has { // user not exist
		return http.StatusUnauthorized, GetUserKeyResponse{"", "incorrect username or password"}
	}
	// ---- get new key ----
	var key = tools.GetKey() // generate new key
	affected, _ := dao.updateUserKey(&User{Key: key}, &User{UserName: username, Password: password})
	if affected == 0 { // user not exist
		fmt.Println("?ssss")
		return http.StatusUnauthorized, GetUserKeyResponse{"", "incorrect username or password"}
	}
	return http.StatusOK, GetUserKeyResponse{key, "get user key successfully"}
}

// GetUserByKeyAndID --- convert string id to int id, if occur error return empty user and error
// check if key is valid and id exsits and belong to the same user
// if valid key and exist id, return User struct
func (*AgendaAtomicService) GetUserByKeyAndID(key string, stringID string) (int, UserKeyResponse) {
	var (
		id   int
		err  error
		has  bool
		user *User
	)
	dao := agendaDao{xormEngine}
	// ---- check key ----
	has, err = dao.ifUserExistByConditions(&User{Key: key})
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

// DeleteUserByKeyAndPassword --- check key if valid
// check if password correct
func (*AgendaAtomicService) DeleteUserByKeyAndPassword(key string, password string) (int, DeleteUserResponse) {
	var (
		err      error
		has      bool
		affected int64
	)
	dao := agendaDao{xormEngine}
	// ---- check key ----
	has, err = dao.ifUserExistByConditions(&User{Key: key})
	if err != nil { // server error
		return http.StatusInternalServerError, DeleteUserResponse{Message: "server error"}
	}
	if !has { // invalid key
		return http.StatusUnauthorized, DeleteUserResponse{Message: "invalid key"}
	}
	// ---- check password ----
	if password == "" { // empty input
		return http.StatusBadRequest, DeleteUserResponse{"empty password"}
	}
	affected, err = dao.deleteUserByKeyAndPassword(key, tools.MD5Encryption(password))
	if err != nil { // server error
		return http.StatusInternalServerError, DeleteUserResponse{Message: "server error"}
	}
	if affected == 0 { // delete user fail
		return http.StatusUnauthorized, DeleteUserResponse{Message: "incorrect paassword"}
	}
	// delete successfully
	return http.StatusNoContent, DeleteUserResponse{}
}

// ListUsersByKeyAndLimit --- check key is valid or not
// if limit is invalid, default set to 10
func (*AgendaAtomicService) ListUsersByKeyAndLimit(key string, stringLimit string) (int, UsersInfoResponse) {
	var (
		limit int

		has   bool
		err   error
		users []User
	)
	dao := agendaDao{xormEngine}
	// ---- check key ----
	has, err = dao.ifUserExistByConditions(&User{Key: key})
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
