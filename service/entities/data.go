package entities

//
// ─── TO BE JSON RESPONSE ───────────────────────────────────────────────────────────
//

// UserKeyResponse -- GetUserByKeyAndID
type UserKeyResponse struct {
	Message  string
	ID       int
	UserName string
	Email    string
	Phone    string
}

// SingleMessageResponse -- DeleteUserByKeyAndPassword/ChangeUserPassword/GetUserKey
type SingleMessageResponse struct {
	Message string
}

// used in UsersInfoResponse
type singleUserInfo struct {
	ID       int
	UserName string
	Email    string
	Phone    string
}

// UsersInfoResponse -- ListUsersByKeyAndLimit
type UsersInfoResponse struct {
	Message            string
	SingleUserInfoList []singleUserInfo
}

// CreateUserResponse -- CreateUser
type CreateUserResponse struct {
	Message  string
	ID       int
	UserName string
	Email    string
	Phone    string
}

// ────────────────────────────────────────────────────────────────────────────────

//
// ─── MESSAGE ────────────────────────────────────────────────────────────────────
//

const (
	// EmptyInput . .
	EmptyInput = "empty input"
	// ServerError .
	ServerError = "server error"
	// DuplicateUsername .
	DuplicateUsername = "duplicate username"
	// CreateUserSuceed .
	CreateUserSuceed = "create user successfully"
	//FailCreateUser .
	FailCreateUser = "fail to create user"
	// EmptyUsernameAndPassword .
	EmptyUsernameAndPassword = "empty username and password"
	// EmptyPassword .
	EmptyPassword = "empty password"
	// IncorrectUsernameAndPassword .
	IncorrectUsernameAndPassword = "incorrect username or password"
	// IncorrectPassword .
	IncorrectPassword = "incorrect password"
	// LoginSucceed .
	LoginSucceed = "login successfully"
	// EmptyID .
	EmptyID = "empty id"
	// InvalidID .
	InvalidID = "invalid id"
	// ReLogin .
	ReLogin = "please re-login"
	// GetUserInfoSucceed .
	GetUserInfoSucceed = "get user info successfully"
	// NotExistedID .
	NotExistedID = "the id does not exist"
	// InvalidLimit .
	InvalidLimit = "invalid limit"
	// InvalidOffset .
	InvalidOffset = "invalid offset"
	// LogoutFail .
	LogoutFail = "log out fail"
	// UpdatePasswordSucceed .
	UpdatePasswordSucceed = "update password successfully"
	// EmptyNewPassword .
	EmptyNewPassword = "new password is empty"
	// NotMatchPassword .
	NotMatchPassword = "new password and confirmation do not match"
)

// ────────────────────────────────────────────────────────────────────────────────
