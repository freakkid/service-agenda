package service

import (
	"net/url"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(createUsername string, createPassword string, createPhone string, createEmail string) error {
	// regist user via http
	resp, err := http.PostForm("http://private-633936-serviceagenda.apiary-mock.com/v1/users?key=1e3576bt", url.Values{"username":{createUsername}, "password":{createPassword}, "phone":{createPhone}, "email":{createEmail}})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// fmt.Println(string(body))
	temp := RetJson{}	
	if err := json.Unmarshal(body, &temp); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return nil
}