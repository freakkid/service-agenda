package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func ListAllUsers(limit string, offset string) (bool, []SingleUserInfo) {
	// list all user via http request
	ok, username, session := GetCurrentUser()
	if !ok {
		fmt.Fprintln(os.Stderr, "Some mistakes happend in ListAllUsers")
		return false, []SingleUserInfo{}
	}
	req, err := http.NewRequest("GET", URL+"/v1/users?limit="+limit+"&offset="+offset, strings.NewReader(""))
	// resp, err := http.Get(URL + "/v1/users?limit=" + limit + "&offset=" + offset)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : some mistakes happend in creating request to server")
		return false, []SingleUserInfo{}
	}
	req.Header.Set("Cookie", username+"="+session)
	fmt.Println(username + "=" + session)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : some mistakes happend in sending request to server")
		return false, []SingleUserInfo{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error : some mistakes happend in forming body")
		return false, []SingleUserInfo{}
	}
	temp := UsersInfoResponse{}
	if err := json.Unmarshal(body, &temp); err != nil {
		fmt.Fprintln(os.Stderr, "error : some mistakes happend in parsing body")
		return false, []SingleUserInfo{}
	}

	// fmt.Println(string(body))
	if resp.StatusCode == 200 {
		ret := make([]SingleUserInfo, len(temp.SingleUserInfoList))
		for index, each := range temp.SingleUserInfoList {
			ret[index].ID = each.ID
			ret[index].UserName = each.UserName
			ret[index].Phone = each.Phone
			ret[index].Email = each.Email
		}
		return false, ret
	}

	fmt.Fprintln(os.Stderr, temp.Message)
	return false, []SingleUserInfo{}
}
