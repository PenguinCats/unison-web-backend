package models

import (
	"fmt"
	"testing"
)

func TestAddMessage(t *testing.T) {
	setupForTest()
	title := "test_title"
	text := "test_text"
	fromUid := int64(1)
	mid, err := AddMessage(title, text, fromUid)
	if err != nil {
		t.Fatal("error when add message")
	}

	message, err := GetMessageByMid(mid)
	if err != nil {
		t.Fatal("error when add message")
	}

	if message.Title != title || message.Text != text || message.FromUid != fromUid {
		t.Fatal("error when add message")
	}
}

func TestGetMessageByMid(t *testing.T) {
	setupForTest()
	message, err := GetMessageByMid(1)
	if err != nil {
		t.Fatal("error")
	}
	fmt.Println(message)
}

func TestGetMessageProfileByMid(t *testing.T) {
	setupForTest()
	messages, err := GetMessageProfileByMids([]int64{1, 2})
	if err != nil {
		t.Fatal("error")
	}
	fmt.Println(messages)
}

func TestGetMessageProfileByUid(t *testing.T) {
	setupForTest()
	messages, err := GetMessageProfileByUid(1)
	if err != nil {
		t.Fatal("error")
	}
	fmt.Println(messages)
}

func TestAddMessageUser(t *testing.T) {
	setupForTest()
	err := AddMessageUser([]int64{1}, 1)
	if err != nil {
		t.Fatal("error")
	}
}

func TestGetMessageUserByUid(t *testing.T) {
	setupForTest()
	messages, err := GetMessageUserByUid(1)
	if err != nil {
		t.Fatal("error")
	}
	fmt.Println(messages)
}

func TestGetMessageProfileByMids(t *testing.T) {
	setupForTest()
	messages, err := GetMessageUserByMids([]int64{1, 2})
	if err != nil {
		t.Fatal("error")
	}
	fmt.Println(messages)
}
