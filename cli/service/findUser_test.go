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
		want RetJson
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindUser(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
