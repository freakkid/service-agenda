package service

import "testing"

func TestGetUserKey(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args {"txz", "1"}, true},
		{"wrong password", args {"txz", "2"}, false},
		{"wrong username", args {"t", ""}, false},
		{"empty", args {"", ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserKey(tt.args.username, tt.args.password); got != tt.want {
				t.Errorf("GetUserKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
