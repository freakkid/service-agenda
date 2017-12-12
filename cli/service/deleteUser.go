package service

import (
	"errors"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
	"os"
)

func DeleteUser(password string) (bool, error) {
	type RetJson struct {
		Status	bool	`json:"status"`
		Message	string	`json:"message"`
	}
	ok, session := GetCurrentUser()
	if !ok {
		return false, errors.New("Some mistakes happend in FindUser")
	}
	url := "https://private-633936-serviceagenda.apiary-mock.com/v1/users/?key="+session+"&password="+password
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return false, errors.New("Can not construct DELETE request.")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, errors.New("Send delete request failed.")
	}

	defer res.Body.Close()
	if res.StatusCode == 204 {
		return true, nil
	} else if res.StatusCode < 500 && res.StatusCode >= 400{
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return false, errors.New("Fail to read body.")
		}
		tmp := RetJson{}
		if err := json.Unmarshal(body, tmp); err != nil {
			fmt.Fprintln(os.Stderr, "Can not resolve body.")
			os.Exit(0)
		}
		return false, errors.New(tmp.Message)
	}
	return false, errors.New("Server failed.")
}