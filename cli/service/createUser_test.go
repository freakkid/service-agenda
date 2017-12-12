package service

import "testing"

func TestCreateUser(t *testing.T) {
	type args struct {
		createUsername string
		createPassword string
		createPhone    string
		createEmail    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.createUsername, tt.args.createPassword, tt.args.createPhone, tt.args.createEmail); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
