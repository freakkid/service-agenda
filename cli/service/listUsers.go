package service

import (
	"fmt"
	"net/http"
	"encoding/json"
	"os"
	"io/ioutil"
)

func ListAllUsers (limit string) []RetJson {
	// list all user via http request
	ok, session := GetCurrentUser()
	if !ok {
		fmt.Fprintln(os.Stderr, "Some mistakes happend in ListAllUsers")
		return []RetJson{}
	}
	resp, err := http.Get("https://private-633936-serviceagenda.apiary-mock.com/v1/users?key="+session+"&limit="+limit)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : some mistakes happend in sending get request to server")
		return []RetJson{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : some mistakes happend in forming body")
		return []RetJson{}
	}
	// fmt.Println(string(body))
	if resp.StatusCode == 200 {
		temp := struct {
			Message				string
			SingleUserInfoList	[]struct {
				Id			int
				Username	string
				Phone		string
				Email		string
			}
		} {}
		if err := json.Unmarshal(body, &temp); err != nil {
			fmt.Fprintln(os.Stderr, "error : some mistakes happend in parsing body")
			return []RetJson{}
		}
		ret := make([]RetJson, len(temp.SingleUserInfoList))
		for index, each := range temp.SingleUserInfoList {
			ret[index].Id = each.Id
			ret[index].Username = each.Username
			ret[index].Phone = each.Phone
			ret[index].Email = each.Email
		}
		return ret
	} else {
		temp := struct {
			Message				string
			SingleUserInfoList	[]struct {
				Id			int
				Username	string
				Phone		string
				Email		string
			}
		} {}
		if err := json.Unmarshal(body, &temp); err != nil {
			fmt.Fprintln(os.Stderr, "error : "+temp.Message)
			return []RetJson{}
		}
		fmt.Fprintln(os.Stderr, temp.Message)
		return []RetJson{}
	}
}