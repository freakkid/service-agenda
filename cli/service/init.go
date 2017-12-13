package service

import (
	"os"
	"os/user"
)

var UserFile 	string
var SessionFile string
var URL string

func init() {
	user, err := user.Current()
	if err !=  nil {
		panic(err)
	}
	UserFile = user.HomeDir+"/currentUser"
	SessionFile = user.HomeDir+"/session"
	envURL := os.Getenv("SERVER_URL")
	PORT := os.Getenv("PORT")
	if len(envURL) == 0 {
		URL = "https://private-633936-serviceagenda.apiary-mock.com"
	} else {
		URL = envURL + PORT
	}
}