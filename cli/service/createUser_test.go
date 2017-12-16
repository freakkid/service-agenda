package service

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"testing"
)

// UserInfo .
type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func TestCreateRes(t *testing.T) {
	// for test
	b, err := json.Marshal(&UserInfo{Username: "hnx", Password: "123", Email: "email@qq.com", Phone: "12345678901"})
	checkErr(err)
	p := ioutil.NopCloser(bytes.NewReader(b))

	type args struct {
		resBody    io.ReadCloser
		statusCode int
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// for test
		{
			name:  "201",
			args:  args{resBody: p, statusCode: 201},
			want:  true,
			want1: "create user successfully",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CreateRes(tt.args.resBody, tt.args.statusCode)
			if got != tt.want {
				t.Errorf("CreateRes() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CreateRes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
