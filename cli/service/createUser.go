package service

import (
	"bytes"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(createUsername string, createPassword string, createPhone string, createEmail string) bool {
	// regist user via http
	// resp, err := http.PostForm("https://private-633936-serviceagenda.apiary-mock.com/v1/users?key="+session+"&limit="+limit, url.Values{"username":{createUsername}, "password":{createPassword}, "phone":{createPhone}, "email":{createEmail}})
	reqBody := fmt.Sprintf("{\"username\": %v, \"password\": %v, \"phone\": %v, \"email\": %v}", createUsername, createPassword, createPhone, createEmail)
	resp, err := http.Post(URL + "/v1/users", "application/json", bytes.NewBufferString(reqBody))
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