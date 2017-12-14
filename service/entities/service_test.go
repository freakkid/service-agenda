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
		want1 UserInfoResponse
	}{
		{
			name:  "create",
			args:  args{username: "hnx", password: "123", email: "email@qq.com", phone: "12345678901"},
			want:  201,
			want1: UserInfoResponse{Message: CreateUserSuceed, ID: 1, UserName: "hnx", Email: "email@qq.com", Phone: "12345678901"},
		},
		{
			name:  "duplicate name",
			args:  args{username: "hnx", password: "123", email: "email@qq.com", phone: "12345678901"},
			want:  400,
			want1: UserInfoResponse{Message: DuplicateUsername, ID: -1, UserName: "", Email: "", Phone: ""},
		},
		{
			name:  "empty input",
			args:  args{username: "", password: "123", email: "email@qq.com", phone: "12345678901"},
			want:  400,
			want1: UserInfoResponse{Message: EmptyInput, ID: -1, UserName: "", Email: "", Phone: ""},
		},
		{name: "successful",
			args:  args{username: "hnx2", password: "123", email: "email@qq.com", phone: "12345678901"},
			want:  201,
			want1: UserInfoResponse{Message: CreateUserSuceed, ID: 2, UserName: "hnx2", Email: "email@qq.com", Phone: "12345678901"},
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
		want1 int
		want2 SingleMessageResponse
	}{
		{
			name:  "login successfully",
			args:  args{username: "hnx", password: "123"},
			want1: 200,
			want2: SingleMessageResponse{LoginSucceed},
		},
		{
			name:  "empty input",
			args:  args{username: "", password: ""},
			want1: 400,
			want2: SingleMessageResponse{EmptyUsernameOrPassword},
		},
		{
			name:  "repeat login",
			args:  args{username: "hnx", password: "123"},
			want1: 200,
			want2: SingleMessageResponse{LoginSucceed},
		},
		{
			name:  "wrong input",
			args:  args{username: "hnx", password: "1??23"},
			want1: 401,
			want2: SingleMessageResponse{IncorrectUsernameAndPassword},
		},
	}
	for _, tt := range tests {
		a := &AgendaAtomicService{}
		_, got1, got2 := a.LoginAndGetSessionID(tt.args.username, tt.args.password)
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
		want1 UserInfoResponse
	}{
		{
			name:  "error sessionid",
			args:  args{sessionID: "hnx", stringID: "1"},
			want:  401,
			want1: UserInfoResponse{Message: ReLogin, ID: -1},
		},
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
		{
			name:  "error sessionid",
			args:  args{sessionID: "hnx", stringID: "1"},
			want:  401,
			want1: SingleMessageResponse{Message: ReLogin},
		},
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
		{
			name:  "error sessionid",
			args:  args{sessionID: "hnx"},
			want:  401,
			want1: UsersInfoResponse{ReLogin, []singleUserInfo{}},
		},
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
		{
			name:  "error sessionid",
			args:  args{sessionID: "hnx", stringID: "1"},
			want:  401,
			want1: SingleMessageResponse{Message: ReLogin},
		},
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
		{
			name:  "error sessionid",
			args:  args{sessionID: "hnx", stringID: "1"},
			want:  401,
			want1: SingleMessageResponse{Message: LogoutFail},
		},
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
