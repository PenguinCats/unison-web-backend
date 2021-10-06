package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	Mid     int64  `gorm:"column:mid; primaryKey"`
	Title   string `gorm:"column:title"`
	Text    string `gorm:"column:text"`
	FromUid int64  `gorm:"column:from_uid"`

	CreatedAt time.Time `gorm:"column:sent_time"`
}

func (Message) TableName() string {
	return "message"
}

func GetMessageByMid(mid int64) (*Message, error) {
	var message Message
	err := db.First(&message, mid).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func GetMessageProfileByMids(mids []int64) (*[]Message, error) {
	var message []Message
	err := db.Select("mid", "title", "from_uid", "sent_time").Find(&message, mids).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func GetMessageProfileByUid(uid int64) (*[]Message, error) {
	var message []Message
	err := db.Select("mid", "title", "from_uid", "sent_time").Where("from_uid = ?", uid).Find(&message).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &message, nil
}

//func AddMessage(title, text string, fromUid int64) (int64, error) {
//	message := Message{
//		Title:   title,
//		Text:    text,
//		FromUid: fromUid,
//	}
//	err := db.Create(&message).Error
//	if err != nil {
//		return 0, err
//	}
//
//	return message.Mid, nil
//}

type MessageUser struct {
	ID    int64 `gorm:"column:id"`
	ToUid int64 `gorm:"column:to_uid"`
	Mid   int64 `gorm:"column:mid"`
	State int64 `gorm:"column:state"`
}

func (MessageUser) TableName() string {
	return "message_user"
}

func GetMessageUserByUid(uid, offset, limit int64) (*[]MessageUser, error) {
	var messageUser []MessageUser
	err := db.Where("to_uid = ?", uid).Not("state = 3").Find(&messageUser).Error
	if err != nil {
		return nil, err
	}
	return &messageUser, nil
}

func GetMessageUserByMids(mids []int64) (*[]MessageUser, error) {
	var messageUser []MessageUser
	err := db.Where("mid IN ?", mids).Find(&messageUser).Error
	if err != nil {
		return nil, err
	}
	return &messageUser, nil
}

//func AddMessageUser(uids []int64, mid int64) error {
//	messageUsers := make([]MessageUser, len(uids))
//	for idx, uid := range uids {
//		messageUsers[idx].ToUid = uid
//		messageUsers[idx].Mid = mid
//		messageUsers[idx].State = 1
//	}
//
//	err := db.Create(&messageUsers).Error
//	return err
//}

func GetMessageNumberByToUid(uid int64) (int64, error) {
	var count int64
	err := db.Model(&MessageUser{}).Where("to_uid = ?", uid).Not("state = 3").Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func CheckMessageUserExist(mid, uid int64) (*MessageUser, error) {
	var mu MessageUser
	return &mu, db.Where(&MessageUser{
		ToUid: uid,
		Mid:   mid,
	}).Not("state = 3").Take(&mu).Error
}

func MarkAsRead(id int64) error {
	err := db.Model(&MessageUser{ID: id}).Update("state", 2).Error
	return err
}

func MarkAsDelete(id int64) error {
	err := db.Model(&MessageUser{ID: id}).Update("state", 3).Error
	return err
}

// union operation

func AddMessageUnion(fromUid int64, toUids []int64, title, text string) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		message := Message{
			Title:   title,
			Text:    text,
			FromUid: fromUid,
		}
		if err := tx.Create(&message).Error; err != nil {
			return err
		}

		messageUsers := make([]MessageUser, len(toUids))
		for idx, uid := range toUids {
			messageUsers[idx].ToUid = uid
			messageUsers[idx].Mid = message.Mid
			messageUsers[idx].State = 1
		}
		if err := tx.Create(&messageUsers).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
