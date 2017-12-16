package service

import (
	"net/http"
	"testing"
)

func TestUpdatePassword(t *testing.T) {
	type args struct {
		old     string
		new     string
		confirm string
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
			got, err := UpdatePassword(tt.args.old, tt.args.new, tt.args.confirm)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdatePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateRes(t *testing.T) {
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
			got, err := UpdateRes(tt.args.res)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateRes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UpdateRes() = %v, want %v", got, tt.want)
			}
		})
	}
}
