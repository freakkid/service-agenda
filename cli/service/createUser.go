package service

import (
	"net/url"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(createUsername string, createPassword string, createPhone string, createEmail string, limit string) bool {
	ok, session := GetCurrentUser()
	if !ok {
		fmt.Fprintln(os.Stderr, "Some mistakes happend in createUser")
		return false
	}
	// regist user via http
	resp, err := http.PostForm("https://private-633936-serviceagenda.apiary-mock.com/v1/users?key="+session+"&limit="+limit, url.Values{"username":{createUsername}, "password":{createPassword}, "phone":{createPhone}, "email":{createEmail}})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : Some mistakes happend in reading resp.Body" )
		return false
	}

	fmt.Println(string(body))
	if resp.StatusCode == 201  {
		temp := struct {
			Status 		bool
			Message		string
		} {}
		if err := json.Unmarshal(body, &temp); err != nil {
			fmt.Fprintln(os.Stderr, "error : Some mistakes happend in parsing resp.Body")
			return false
		}
		return true
	} else {
		temp := struct {
			Message		string
			Id			int
			Username	string
			Phone 		string
			Email 		string
		} {}
		if err := json.Unmarshal(body, &temp); err != nil {
			fmt.Fprintln(os.Stderr, "error : Some mistakes happend in unmarshal")
			return false
		}
		fmt.Fprintln(os.Stderr, temp.Message)
		return false
	}
}