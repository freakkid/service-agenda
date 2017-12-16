package service

import (
	"net/http"
	"testing"
)

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
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateUser(tt.args.createUsername, tt.args.createPassword, tt.args.createPhone, tt.args.createEmail); got != tt.want {
				t.Errorf("CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateRes(t *testing.T) {
	type args struct {
		resp *http.Response
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
			if got := CreateRes(tt.args.resp); got != tt.want {
				t.Errorf("CreateRes() = %v, want %v", got, tt.want)
			}
		})
	}
}
