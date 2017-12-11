package entities

import "github.com/go-xorm/xorm"

type agendaDao struct {
	*xorm.Engine
}

//
// ─── OPERATIONS ON USERS INTO DATABASE ───────────────────────────────────────────
//

// create user into database
func (dao *agendaDao) createUser(user *User) error {
	_, err := dao.Insert(user)
	return err
}

func (dao *agendaDao) findUserByUID(id int) (*User, error) {
	var user = &User{ID: id}
	_, err := dao.Get(user)
	return user, err
}

// get a user by user name and password
func (dao *agendaDao) findUserByUsernamePassword(username string, password string) (*User, error) {
	var user = &User{UserName: username, Password: password}
	_, err := dao.Get(user)
	return user, err
}
func (dao *agendaDao) UpdateKeyByUsernamePassword(username string, password string, key string) (*User, error) {
	var user = &User{UserName: username, Password: password}
	_, err := dao.Get(user)
	return user, err
}

// get all users info
func (dao *agendaDao) getLimitUsers(limitNumber int) ([]User, error) {
	if limitNumber <= 0 {
		limitNumber = 20
	}
	var userList = make([]User, 0, 0)
	err := dao.Limit(limitNumber).Find(&userList)
	return userList, err
}

// count all users
func (dao *agendaDao) countAllUsers() (int64, error) {
	return dao.Count(new(User))
}

// delete user by id
func (dao *agendaDao) deleteUserByID(id int) (int64, error) {
	return dao.Delete(&User{ID: id})
}

//
// ─── OPERATIONS ON MEETINGS INTO DATABASE ───────────────────────────────────────────
//

// create meeting into database
func (dao *agendaDao) createMeeting(meeting *Meeting) error {
	_, err := dao.Insert(meeting)
	return err
}

// get a meeting by title
func (dao *agendaDao) findMeetingByTitle(title string) (*Meeting, error) {
	var meeting = &Meeting{Title: title}
	_, err := dao.Get(meeting)
	return meeting, err
}

// get meeting list that user sponsored
func (dao *agendaDao) findMeetingBySponsor(sponsor string) ([]Meeting, error) {
	var meetingList = make([]Meeting, 0, 0)
	err := dao.In("sponsor", sponsor).Find(&meetingList)
	return meetingList, err
}

// get meeting list that user sponsored or participated in by start time and end time
func (dao *agendaDao) countMeetingsByTimeAndSponsor(sponsor string, startTime string, endTime string) (int64, error) {
	return dao.Where("sponsor = ?", sponsor).
		Where("AND start_time <= ? AND end_time >= ? OR start_time <= ? AND end_time >= ? OR start_time >= ? AND end_time <= ?",
			startTime, startTime, endTime, endTime, startTime, endTime).Count()
}

func (dao *agendaDao) countMeetingsByTimeAndTitle(title string, startTime string, endTime string) (int64, error) {
	return dao.Where("title = ?", title).
		Where("AND start_time <= ? AND end_time >= ? OR start_time <= ? AND end_time >= ? OR start_time >= ? AND end_time <= ?",
			startTime, startTime, endTime, endTime, startTime, endTime).Count()
}

// get conflict meetings that user sponsored or participated in by start time and end time
func (dao *agendaDao) countConflictMeetingsByTimeAndSponsor(
	sponsor string, startTime string, endTime string) (int64, error) {
	return dao.Where("sponsor = ?", sponsor).
		Where("AND start_time <= ? AND end_time > ? OR start_time < ? AND end_time >= ? OR start_time >= ? AND end_time <= ?",
			startTime, startTime, endTime, endTime, startTime, endTime).Count()
}

// get conflict meetings that user sponsored or participated in by start time and end time
func (dao *agendaDao) countConflictMeetingsByTimeAndTitle(
	title string, startTime string, endTime string) (int64, error) {
	return dao.Where("title = ?", title).
		Where("AND start_time <= ? AND end_time > ? OR start_time < ? AND end_time >= ? OR start_time >= ? AND end_time <= ?",
			startTime, startTime, endTime, endTime, startTime, endTime).Count()
}

// delete meeting by title
func (dao *agendaDao) deleteMeetingByTitle(title string) (int64, error) {
	return dao.Delete(&Meeting{Title: title})
}

// delete meeting by sponsor name
func (dao *agendaDao) deleteMeetingBySponsor(sponsor string) (int64, error) {
	return dao.Delete(&Meeting{Sponsor: sponsor})
}

//
// ─── OPERATIONS ON PARTICIPATORS INTO DATABASE ───────────────────────────────────────────
//
func (dao *agendaDao) addParticipators(participators *Participators) error {
	_, err := dao.Insert(participators)
	return err
}

func (dao *agendaDao) findParticipatorsByTitle(title string) ([]Participators, error) {
	var participatorsList = make([]Participators, 0, 0)
	err := dao.Find(&participatorsList, &Participators{Title: title})
	return participatorsList, err
}

func (dao *agendaDao) findTitlesByParticipator(participator string) ([]Participators, error) {
	var participatorsList = make([]Participators, 0, 0)
	err := dao.Find(&participatorsList, &Participators{Participator: participator})
	return participatorsList, err
}

func (dao *agendaDao) countParticipatorsByTitle(title string) (int64, error) {
	return dao.Where("title = ?", title).Count()
}

func (dao *agendaDao) deleteParticipatorsByTitle(title string) (int64, error) {
	return dao.Delete(&Participators{Title: title})
}

func (dao *agendaDao) deleteTitlesByParticipator(participator string) (int64, error) {
	return dao.Delete(&Participators{Participator: participator})
}
