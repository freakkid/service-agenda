package service

import "testing"

// {
// 	"username":"zhang3",
// 	"password":"seerfrgfg",
// 	"phone":"12345678901",
// 	"email":"zhang3@mail2.sysu.edu.cn",
// }
func TestCreateUser(t *testing.T) {
	type args struct {
		createUsername string
		createPassword string
		createPhone    string
		createEmail    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"2", args {"", "pass", "12345678901", "1@1.com"}, false},
		{"2", args {"1", "", "12345678901", "1@1.com"}, false},
		{"2", args {"1", "pass", "1", "1@1.com"}, false},
		{"2", args {"1", "pass", "12345678901", "11.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateUser(tt.args.createUsername, tt.args.createPassword, tt.args.createPhone, tt.args.createEmail); got != tt.want {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
