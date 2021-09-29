package models

import (
	"fmt"
	"testing"
)

func TestGetUIDByUserName(t *testing.T) {
	setupForTest()
	uid, err := GetUIDByUserName("zhangbinjie")
	fmt.Println(uid, err)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestExistUserByUserName(t *testing.T) {
	setupForTest()
	flag, err := ExistUserByUserName("zhangbinjie")
	fmt.Println(flag, err)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetUserByUID(t *testing.T) {
	setupForTest()
	user, err := GetUserByUID(1)
	fmt.Println(user, err)
	if err != nil {
		t.Fatal(err.Error())
	}

	user, err = GetUserByUID(0)
	fmt.Println(user, err)
	if err == nil {
		t.Fatal("查询到不存在数据")
	}
}

func TestGetUserAll(t *testing.T) {
	setupForTest()
	users, err := GetUserAll()
	for _, user := range users {
		fmt.Println(*user)
	}
	println(err)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetPwdByUid(t *testing.T) {
	setupForTest()
	pwd, err := GetPwdByUid(1)
	fmt.Println(pwd, err)
	if err != nil {
		t.Fatal(err.Error())
	}
	if pwd != "35468eae9c9d874927ae7f13991e437a" {
		t.Fatal("查询到错误数据")
	}
}
