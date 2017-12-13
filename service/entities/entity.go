package entities

import (
	"time"
)

// User is an entity to save user info with username is primary key.
type User struct {
	ID       int    `xorm:"pk autoincr"`
	Key      string `xorm:"varchar(255) unique"`
	UserName string `xorm:"varchar(255) notnull unique"`
	Password string `xorm:"varchar(255) notnull"`
	Email    string `xorm:"varchar(255) notnull"`
	Phone    string `xorm:"varchar(255) notnull"`
}

// Meeting is an entity to save meeting info with title is primary key,
// names of particulars will created by join string array with space,
// the number of particulars must larger than zero,
// time is used "*time.Time" and save in database as "YYYY-MM-DD HH:mm:SS".
type Meeting struct {
	Title     string     `xorm:"pk varchar(255) notnull unique"`
	Sponsor   string     `xorm:"varchar(255) notnull"`
	StartTime *time.Time `xorm:"DateTime notnull"`
	EndTime   *time.Time `xorm:"DateTime notnull"`
}

// Participators is Participator of meeting select by title
type Participators struct {
	Title        string `xorm:"varchar(255) notnull"`
	Participator string `xorm:"varchar(255) notnull"`
}

// MeetingParticipators to join data of two tables
type MeetingParticipators struct {
	Meeting       `xorm:"extends"`
	Participators `xorm:"extends"`
}

// UserMeeting to join data of two tables
type UserMeeting struct {
	User    `xorm:"extends"`
	Meeting `xorm:"extends"`
}
