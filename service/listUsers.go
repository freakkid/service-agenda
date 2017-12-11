package service

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"io/ioutil"
)

func ListAllUsers () []RetJson {
	// list all user via http request
	ok, session := GetCurrentUser()
	if !ok {
		fmt.Fprintln(os.Stderr, "Some mistakes happend in ListAllUsers")
	}
	resp, err := http.Get("https://private-633936-serviceagenda.apiary-mock.com/v1/users?key="+session)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : some mistakes happend in sending get request to server")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : some mistakes happend in forming body")
	}
	// fmt.Println(string(body))
	var temp []RetJson
	if err := json.Unmarshal(body, &temp); err != nil {
		fmt.Fprintln(os.Stderr, "error : some mistakes happend in parsing body")
	}
	return temp
}