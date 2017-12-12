package service

import "testing"

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
		{"normal", args {"txz", "1", "11111111111", "1@1.com"}, true},
		{"wrong", args {"", "", "", ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateUser(tt.args.createUsername, tt.args.createPassword, tt.args.createPhone, tt.args.createEmail); got != tt.want {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
