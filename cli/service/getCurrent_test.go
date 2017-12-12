package service

import "testing"

func TestGetCurrentUser(t *testing.T) {
	tests := []struct {
		name  string
		want  bool
		want1 string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetCurrentUser()
			if got != tt.want {
				t.Errorf("GetCurrentUser() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetCurrentUser() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
