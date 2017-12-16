package service

import (
	"net/http"
	"testing"
)

func TestLogout(t *testing.T) {
	tests := []struct {
		name    string
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Logout()
			if (err != nil) != tt.wantErr {
				t.Errorf("Logout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Logout() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogoutRes(t *testing.T) {
	type args struct {
		res *http.Response
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LogoutRes(tt.args.res)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogoutRes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("LogoutRes() = %v, want %v", got, tt.want)
			}
		})
	}
}
