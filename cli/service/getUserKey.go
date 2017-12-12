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
			os.Exit(1)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error ： Some mistakes happend in forming body")
			os.Exit(1)
		}
		RetJson := struct {
			Key			string
			Permissions	[]string
		} {}
		if err = json.Unmarshal(body, &RetJson); err != nil {
			fmt.Fprintln(os.Stderr, "error: some mistakes happend in parsing body")
			os.Exit(1)
		} 
		// fmt.Println(string(body))
		if RetJson.Key == "" {
			if times < 2 {
				times++
				fmt.Print("Wrong password, Please try again: ")
				fmt.Scanf("%s", &password)
			} else {
				fmt.Fprintln(os.Stderr, "error : Wrong password")
				os.Exit(1)
			}
		} else {
			// write to file -- session
			err := ioutil.WriteFile(SessionFile, []byte(RetJson.Key), 0655) 
			if err != nil {
				fmt.Fprintln(os.Stderr, "Some mistakes happend in write to session")
				panic(err)
			}
			// write to file -- current
			if ioutil.WriteFile(UserFile, []byte(username), 0655) != nil {
				fmt.Fprintln(os.Stderr, "Some mistakes happend in write to currentUser")
				os.Exit(1)
			}
			break
		}
	}
	return true
}