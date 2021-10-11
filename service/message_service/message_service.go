package message_service

import (
	"errors"
	"github.com/PenguinCats/unison-web-backend/models"
	"github.com/PenguinCats/unison-web-backend/pkg/e"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID      int64
	Mid     int64
	Title   string
	Text    string
	FromUid int64
	ToUid   int64
	ToUids  []int64
	State   int64

	PageSize   int64
	PageNumber int64

	CreatedAt time.Time `gorm:"column:sent_time"`
}

func (m *Message) AddMessage() error {
	return models.AddMessageUnion(m.FromUid, m.ToUids, m.Title, m.Text)
}

func (m *Message) GetMessageInboxProfileByToUid() (*[]models.Message, *[]models.MessageUser, error) {
	offset := m.PageSize * (m.PageNumber - 1)
	if offset < 0 {
		offset = 0
	}
	if m.PageSize <= 0 {
		m.PageSize = 10
	}
	messageUserList, err := models.GetMessageUserByUid(m.ToUid, offset, m.PageSize)
	if err != nil {
		return nil, nil, err
	}

	var mids []int64
	for _, item := range *messageUserList {
		mids = append(mids, item.Mid)
	}

	messageList, err := models.GetMessageProfileByMids(mids)
	if err != nil {
		return nil, nil, err
	}
	mp := make(map[int64]models.Message)
	for _, item := range *messageList {
		mp[item.Mid] = item
	}

	var messageListRet []models.Message
	for _, item := range *messageUserList {
		val, ok := mp[item.Mid]
		if !ok {
			return nil, nil, errors.New("message broken")
		}
		messageListRet = append(messageListRet, val)
	}

	return &messageListRet, messageUserList, nil
}

func (m *Message) GetTotalMessageCountByToUid() (int64, error) {
	return models.GetMessageNumberByToUid(m.ToUid)
}

func (m *Message) GetMessageInboxDetail() int {
	_, err := models.CheckMessageUserExist(m.Mid, m.ToUid)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return e.ERROR_AUTH_PERMISSION_DENIED
		}
		return e.ERROR
	}

	message, err := models.GetMessageByMid(m.Mid)
	if err != nil {
		return e.ERROR
	}

	m.Title = message.Title
	m.Text = message.Text
	m.FromUid = message.FromUid
	m.CreatedAt = message.CreatedAt

	return e.SUCCESS
}

func (m *Message) CheckIsReceiver() int {
	messageUser, err := models.CheckMessageUserExist(m.Mid, m.ToUid)
	if err == gorm.ErrRecordNotFound {
		return e.ERROR_AUTH_PERMISSION_DENIED
	}

	if err != nil {
		return e.ERROR
	}

	m.ID = messageUser.ID
	return e.SUCCESS
}

func (m *Message) MarkAsUserRead() int {
	err := models.MarkAsRead(m.ID)
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

func (m *Message) MarkAsUserDelete() int {
	err := models.MarkAsDelete(m.ID)
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}
