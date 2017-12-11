package service

import (
	"os/user"
)

var UserFile 	string
var SessionFile string

func init() {
	user, err := user.Current()
	if err !=  nil {
		panic(err)
	}
	UserFile = user.HomeDir+"/currentUser"
	SessionFile = user.HomeDir+"/session"
}