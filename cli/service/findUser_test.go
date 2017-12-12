package service

import (
	"reflect"
	"testing"
)

func TestFindUser(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args {"1"}, true},
		{"wrong", args {""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := FindUser(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
