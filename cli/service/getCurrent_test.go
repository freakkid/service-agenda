package service

import "testing"

func TestGetCurrentUser(t *testing.T) {
	got, got1 := GetCurrentUser()
	if got && len(got1) == 0 {
		t.Errorf("GetCurrentUser() get empty answer")
	}
}
