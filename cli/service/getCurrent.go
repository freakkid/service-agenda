package service

import (
	"io/ioutil"
)

func GetCurrentUser() (bool, string) {
	name , err := ioutil.ReadFile(UserFile)	
	if err != nil {
		return false, "Some mistakes happend in read currentUserName from 'current-user'"
	}
	return true, string(name)
}