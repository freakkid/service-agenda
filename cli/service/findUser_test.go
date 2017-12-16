package service

import (
	"net/http"
	"reflect"
	"testing"
)

func TestFindUser(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 SingleUserInfo
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindUser(tt.args.id)
			if got != tt.want {
				t.Errorf("FindUser() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FindUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFindRes(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 SingleUserInfo
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindRes(tt.args.resp)
			if got != tt.want {
				t.Errorf("FindRes() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("FindRes() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
