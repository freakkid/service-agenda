package entities

import (
	"reflect"
	"testing"
)

func TestAgendaAtomicService_CreateUser(t *testing.T) {
	type args struct {
		username string
		password string
		email    string
		phone    string
	}
	tests := []struct {
		name  string
		a     *AgendaAtomicService
		args  args
		want  int
		want1 CreateUserResponse
	}{
		{
			name:  "fgh",
			a:     &AgendaService,
			args:  args{username: "hnx", password: "123", email: "email@qq.com", phone: "12345678901"},
			want:  200,
			want1: CreateUserResponse{},
		},
	}
	for _, tt := range tests {
		a := &AgendaAtomicService{}
		got, got1 := a.CreateUser(tt.args.username, tt.args.password, tt.args.email, tt.args.phone)
		if got != tt.want {
			t.Errorf("%q. AgendaAtomicService.CreateUser() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. AgendaAtomicService.CreateUser() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestAgendaAtomicService_LoginAndGetSessionID(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name  string
		a     *AgendaAtomicService
		args  args
		want  string
		want1 int
		want2 SingleMessageResponse
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &AgendaAtomicService{}
		got, got1, got2 := a.LoginAndGetSessionID(tt.args.username, tt.args.password)
		if got != tt.want {
			t.Errorf("%q. AgendaAtomicService.LoginAndGetSessionID() got = %v, want %v", tt.name, got, tt.want)
		}
		if got1 != tt.want1 {
			t.Errorf("%q. AgendaAtomicService.LoginAndGetSessionID() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
		if !reflect.DeepEqual(got2, tt.want2) {
			t.Errorf("%q. AgendaAtomicService.LoginAndGetSessionID() got2 = %v, want %v", tt.name, got2, tt.want2)
		}
	}
}

func TestAgendaAtomicService_GetUserInfoByID(t *testing.T) {
	type args struct {
		sessionID string
		stringID  string
	}
	tests := []struct {
		name  string
		a     *AgendaAtomicService
		args  args
		want  int
		want1 UserKeyResponse
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &AgendaAtomicService{}
		got, got1 := a.GetUserInfoByID(tt.args.sessionID, tt.args.stringID)
		if got != tt.want {
			t.Errorf("%q. AgendaAtomicService.GetUserInfoByID() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. AgendaAtomicService.GetUserInfoByID() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestAgendaAtomicService_DeleteUserByPassword(t *testing.T) {
	type args struct {
		sessionID string
		stringID  string
		password  string
	}
	tests := []struct {
		name  string
		a     *AgendaAtomicService
		args  args
		want  int
		want1 SingleMessageResponse
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &AgendaAtomicService{}
		got, got1 := a.DeleteUserByPassword(tt.args.sessionID, tt.args.stringID, tt.args.password)
		if got != tt.want {
			t.Errorf("%q. AgendaAtomicService.DeleteUserByPassword() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. AgendaAtomicService.DeleteUserByPassword() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestAgendaAtomicService_ListUsersByLimit(t *testing.T) {
	type args struct {
		sessionID    string
		stringLimit  string
		stringOffset string
	}
	tests := []struct {
		name  string
		a     *AgendaAtomicService
		args  args
		want  int
		want1 UsersInfoResponse
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &AgendaAtomicService{}
		got, got1 := a.ListUsersByLimit(tt.args.sessionID, tt.args.stringLimit, tt.args.stringOffset)
		if got != tt.want {
			t.Errorf("%q. AgendaAtomicService.ListUsersByLimit() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. AgendaAtomicService.ListUsersByLimit() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestAgendaAtomicService_ChangeUserPassword(t *testing.T) {
	type args struct {
		sessionID    string
		stringID     string
		password     string
		newPassword  string
		confirmation string
	}
	tests := []struct {
		name  string
		a     *AgendaAtomicService
		args  args
		want  int
		want1 SingleMessageResponse
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &AgendaAtomicService{}
		got, got1 := a.ChangeUserPassword(tt.args.sessionID, tt.args.stringID, tt.args.password, tt.args.newPassword, tt.args.confirmation)
		if got != tt.want {
			t.Errorf("%q. AgendaAtomicService.ChangeUserPassword() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. AgendaAtomicService.ChangeUserPassword() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestAgendaAtomicService_LogoutAndDeleteSessionID(t *testing.T) {
	type args struct {
		sessionID string
		stringID  string
	}
	tests := []struct {
		name  string
		a     *AgendaAtomicService
		args  args
		want  int
		want1 SingleMessageResponse
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		a := &AgendaAtomicService{}
		got, got1 := a.LogoutAndDeleteSessionID(tt.args.sessionID, tt.args.stringID)
		if got != tt.want {
			t.Errorf("%q. AgendaAtomicService.LogoutAndDeleteSessionID() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. AgendaAtomicService.LogoutAndDeleteSessionID() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}
