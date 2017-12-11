package service

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
	"io/ioutil"
)

type RetJson struct {
	Id			int		`json:"id"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Phone		string	`json:"phone"`
	Email		string	`json:"email"`
}

func FindUser(id string) RetJson {
	ok, session := GetCurrentUser()
	if !ok {
		fmt.Fprintln(os.Stderr, "Some mistakes happend in FindUser")
		os.Exit(1)
	}
	tarUrl := "https://private-633936-serviceagenda.apiary-mock.com/v1/users/?key="+session+"&id="+id
	resp, err := http.Get(tarUrl)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : Some mistakes happend in sending get request to tarUrl")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : Some mistakes happend in forming body")
	}
	// fmt.Println(string(body))
	temp := RetJson{}
	if err := json.Unmarshal(body, &temp); err != nil {
		fmt.Fprintln(os.Stderr, "error : Some mistakes happend in parsing body")
	}
	return temp
}