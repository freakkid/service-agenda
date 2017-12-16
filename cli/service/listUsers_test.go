package service

import (
	"net/http"
	"reflect"
	"testing"
)

func TestListAllUsers(t *testing.T) {
	type args struct {
		limit  string
		offset string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []SingleUserInfo
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ListAllUsers(tt.args.limit, tt.args.offset)
			if got != tt.want {
				t.Errorf("ListAllUsers() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListAllUsers() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestListRes(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []SingleUserInfo
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ListRes(tt.args.resp)
			if got != tt.want {
				t.Errorf("ListRes() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ListRes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
