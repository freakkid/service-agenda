package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetUserKey(username string, password string) bool {
	var times int
	type RetJson struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	}
	for {
		tarUrl := URL + "/v1/user/login?username=" + username + "&password=" + password
		resp, err := http.Get(tarUrl)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error : Some mistakes happend in sending get request to tarUrl")
			return false
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error ： Some mistakes happend in forming body")
			return false
		}
		temp := RetJson{}
		if err = json.Unmarshal(body, &temp); err != nil {
			fmt.Fprintln(os.Stderr, "error: some mistakes happend in parsing body")
			return false
		}
		// fmt.Println(string(body))
		if resp.StatusCode != 200 {
			fmt.Println(temp.Message)
			if times < 2 {
				times++
				fmt.Print("Wrong password, Please try again: ")
				fmt.Scanf("%s", &password)
			} else {
				fmt.Fprintln(os.Stderr, "error : Wrong password")
				return false
			}
		} else {
			session := ""
			for _, item := range resp.Cookies() {
				if item.Name == "key" {
					session = item.Value
				}
			}
			if session == "" {
				fmt.Fprintln(os.Stderr, "error : session should not be empty")
				return false
			}
			fmt.Println("Geted session : " + session)
			// write to file -- user
			err = ioutil.WriteFile(UserFile, []byte(temp.ID), 0655)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Some mistakes happend in writing to current user")
			}
			// write to file -- session
			err = ioutil.WriteFile(SessionFile, []byte(session), 0655)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Some mistakes happend in writing to session")
				return false
			}
			break
		}
	}
	return true
}
