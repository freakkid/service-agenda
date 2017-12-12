package service 

import (
	"io/ioutil"
	"os"
	"fmt"
	"net/http"
	"encoding/json"
)

func GetUserKey(username string, password string) bool {
	var times int
	for {
		tarUrl := "https://private-633936-serviceagenda.apiary-mock.com/v1/user/getkey?username=" + username + "&password=" +password
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
		temp := struct {
			Key			string
			Status		bool
			Message		string
		} {}
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
			// write to file -- session
			err := ioutil.WriteFile(SessionFile, []byte(temp.Key), 0655) 
			if err != nil {
				fmt.Fprintln(os.Stderr, "Some mistakes happend in write to session")
				return false
			}
			break
		}
	}
	return true
}