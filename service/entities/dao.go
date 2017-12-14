package entities

import "github.com/go-xorm/xorm"

type agendaDao struct {
	*xorm.Engine
}

//
// ─── OPERATIONS ON USERS INTO DATABASE ───────────────────────────────────────────
//

// create user into database -- return result and id of user
func (dao *agendaDao) createUser(user *User) (bool, *User) {
	affected, _ := dao.Insert(user)
	if affected == 1 {
		return true, user
	}
	return false, nil
}

func (dao *agendaDao) updateUser(user *User, selectedUser *User) (int64, error) {
	return dao.Update(user, selectedUser)
}

func (dao *agendaDao) ifUserExistByConditions(user *User) (bool, error) {
	return dao.Get(user)
}

func (dao *agendaDao) findUserByConditions(user *User) (bool, *User) {
	has, err := dao.ifUserExistByConditions(user)
	if has && err == nil {
		return has, user
	}
	return has, nil
}

// get all users info
func (dao *agendaDao) getLimitUsers(limitNumber int, offsetNumber int) ([]User, error) {
	if limitNumber <= 0 {
		limitNumber = 5
	}
	if offsetNumber < 0 {
		offsetNumber = 0
	}
	var userList = make([]User, 0, 0)
	err := dao.Limit(limitNumber, offsetNumber).Find(&userList)
	return userList, err
}

// count all users
func (dao *agendaDao) countAllUsers() (int64, error) {
	return dao.Count(new(User))
}

// delete user by sessionID and password
func (dao *agendaDao) deleteUserBySessionIDAndIDAndPassword(sessionID string, id int, password string) (int64, error) {
	return dao.Delete(&User{SessionID: sessionID, ID: id, Password: password})
}

//+++++++++++++++++++++++++++++The funtions below have not been used+++++++++++++++++++++++++++++++++++++++++++++++++
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
