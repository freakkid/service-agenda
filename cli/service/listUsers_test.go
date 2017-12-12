package service

import (
	"reflect"
	"testing"
)

func TestListAllUsers(t *testing.T) {
	tests := []struct {
		name string
		want []RetJson
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListAllUsers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAllUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
