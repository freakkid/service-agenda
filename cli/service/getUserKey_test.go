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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUserKey(tt.args.username, tt.args.password); got != tt.want {
				t.Errorf("GetUserKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
