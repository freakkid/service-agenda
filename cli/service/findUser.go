package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type RetJson struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func FindUser(id string) (bool, SingleUserInfo) {
	ok, username, session := GetCurrentUser()
	if !ok {
		fmt.Fprintln(os.Stderr, "Some mistakes happend in FindUser")
		return false, SingleUserInfo{}
	}
	tarUrl := URL + "/v1/users/?id=" + session + "&id=" + id
	req, err := http.NewRequest("GET", tarUrl, strings.NewReader(""))
	// resp, err := http.Get(URL + "/v1/users?limit=" + limit + "&offset=" + offset)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : some mistakes happend in creating request to server")
		return false, SingleUserInfo{}
	}
	req.Header.Set("Cookie", "key="+session)
	fmt.Println(username + "=" + session)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : some mistakes happend in sending request to server")
		return false, SingleUserInfo{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : Some mistakes happend in forming body")
	}
	// fmt.Println(string(body))
	if resp.StatusCode == 200 {
		temp := UserKeyResponse{}
		if err := json.Unmarshal(body, &temp); err != nil {
			fmt.Fprintln(os.Stderr, "error : Some mistakes happend in parsing body")
		}
		return true, SingleUserInfo{temp.ID, temp.UserName, temp.Email, temp.Phone}
	} else {
		temp := UserKeyResponse{}
		if err := json.Unmarshal(body, &temp); err != nil {
			fmt.Fprintln(os.Stderr, "error : Some mistakes happend in parsing body")
		}
		fmt.Fprintln(os.Stderr, temp.Message)
		return false, SingleUserInfo{}
	}
}
