package service

import (
	"os"
)

var UserFile string
var SessionFile string
var UserMap string
var URL string

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
type SingleUserInfo struct {
	ID       int
	UserName string
	Email    string
	Phone    string
}

// UsersInfoResponse -- ListUsersByKeyAndLimit
type UsersInfoResponse struct {
	Message            string
	SingleUserInfoList []SingleUserInfo
}

// CreateUserResponse -- CreateUser
type CreateUserResponse struct {
	Message  string
	ID       int
	UserName string
	Email    string
	Phone    string
}

type MessageJson struct {
	Message string
}

func init() {
	UserFile = "../currentUser"
	SessionFile = "../session"
	// UserMap = "./userMap"
	envURL := os.Getenv("SERVER_URL")
	PORT := os.Getenv("PORT")
	if len(envURL) == 0 {
		// URL = "https://private-633936-serviceagenda.apiary-mock.com"
		URL = "http://localhost:8080"
	} else {
		URL = envURL + PORT
	}
}
