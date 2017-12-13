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
		{"1", args {""}, false},
		{"2", args {"1q"}, false},
		{"3", args {"qq"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := FindUser(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
