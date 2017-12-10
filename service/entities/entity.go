package entities

import (
	"time"
)

// User is an entity to save user info with username is primary key.
type User struct {
	UserName string `xorm:"pk varchar(255) notnull unique"`
	Password string `xorm:"varchar(255) notnull"`
	Email    string `xorm:"varchar(255) notnull"`
	Phone    string `xorm:"varchar(255) notnull"`
}

// Meeting is an entity to save meeting info with title is primary key,
// names of particulars will created by join string array with space,
// the number of particulars must larger than zero,
// time is used "*time.Time" and save in database as "YYYY-MM-DD HH:mm:SS".
type Meeting struct {
	Title       string
	Sponsor     string
	Particulars string
	StartTime   *time.Time
	EndTime     *time.Time
}
