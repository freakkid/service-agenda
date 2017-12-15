package service

import (
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"errors"
)

func UpdatePassword(old, new, confirm string) (bool, error){
	type RetJson struct {
		Message	string	`json:"message"`
	}
	ok, name, session := GetCurrentUser()
	if !ok {
		return false, errors.New("Some mistakes happend in FindUser")
	}
	url := URL + "/v1/users/" + name + "?password=" + old + "&newpassword=" + new + "&confirmation=" + confirm
	req, err := http.NewRequest("PATCH", url, nil)
	req.Header.Set("cookie", name + "=" + session)
	if err != nil {
		return false, errors.New("Can not construct PATCH request.")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, errors.New("Send patch request failed.")
	}

	defer res.Body.Close()
	if res.StatusCode == 200 {
		return true, nil
	} else if res.StatusCode < 500 && res.StatusCode >= 400 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return false, errors.New("Fail to read body.")
		}
		tmp := RetJson{}
		if err := json.Unmarshal(body, &tmp); err != nil {
			fmt.Fprintln(os.Stderr, "Can not resolve body.")
		}
		return false, errors.New(tmp.Message)
	}
	return false, errors.New("Server failed.")
}