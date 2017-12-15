package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func CreateUser(createUsername string, createPassword string, createPhone string, createEmail string) bool {
	// regist user via http json
	reqBody := fmt.Sprintf("{\"username\": \"%v\", \"password\": \"%v\", \"phone\": \"%v\", \"email\": \"%v\"}", createUsername, createPassword, createPhone, createEmail)
	resp, err := http.Post(URL+"/v1/users", "application/json", bytes.NewBufferString(reqBody))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : Some mistakes happend in reading resp.Body")
		return false
	}

	temp := CreateUserResponse{}
	if err = json.Unmarshal(body, &temp); err != nil {
		fmt.Fprintln(os.Stderr, "error : Some mistakes happend in parsing resp.Body")
		return false
	}
	if resp.StatusCode == 201 {
		// write mapping username ---> id into file UserMap
		f, err := os.OpenFile(UserMap, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0655)
		defer f.Close()
		if err != nil {
			fmt.Println(os.Stderr, "Some mistakes happend in opening UserMap")
			return false
		}
		f.Write([]byte(temp.UserName + "&" + strconv.Itoa(temp.ID) + "\n"))
		return true
	} else {
		return false
	}
}
