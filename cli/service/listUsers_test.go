package service

import (
	"testing"
)

func TestListAllUsers(t *testing.T) {
	type args struct {
		limit string
	}
	tests := []struct {
		name string
		args args
		want []RetJson
	}{
		{"normal", args {"5"}, []RetJson{}},
		{"wrong limit1", args {"-1"}, []RetJson{}},
		{"wrong limit2", args {"2"}, []RetJson{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ListAllUsers(tt.args.limit)
			if len(got) > 5 {
				t.Errorf("ListUsers() want at most %v, but get %v", tt.args.limit, len(got))
			}
		})
	}
}
