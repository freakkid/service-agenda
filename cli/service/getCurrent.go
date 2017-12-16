package service

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func GetCurrentUser() (bool, string, string) {
	UserItem, err := ioutil.ReadFile(UserFile)
	if err != nil {
		fmt.Println(os.Stderr, "error : Some mistakes happend in reading UserFile")
		return false, "", ""
	}
	SessionItem, err := ioutil.ReadFile(SessionFile)
	if err != nil {
		fmt.Println(os.Stderr, "error : Some mistakes happend in reading SessionFile")
		return false, "", ""
	}
	username := string(UserItem)
	session := string(SessionItem)
	if err != nil || len(username) == 0 || len(session) == 0 {
		return false, "", ""
	}
	return true, username, session
}

func Username2Id(username string) (bool, string) {
	f, err := os.Open(UserMap)
	if err != nil {
		fmt.Println(os.Stderr, "error : Some mistakes happend in opening UserMap")
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
		if strings.Split(line, "&")[0] == username {
			return true, strings.Split(line, "&")[1]
		}
	}
	return false, ""
}

func RemoveFile() {
	os.Remove(UserMap)
	os.Remove(UserFile)
	os.Remove(SessionFile)
}
