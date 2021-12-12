package models

import (
	"fmt"
	"testing"
)

func TestGetUIDByUserName(t *testing.T) {
	SetupForTest()
	uid, err := GetUIDByUserName("zhangbinjie")
	fmt.Println(uid, err)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestExistUserByUserName(t *testing.T) {
	SetupForTest()
	flag, err := ExistUserByUserName("zhangbinjie")
	fmt.Println(flag, err)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetUserByUID(t *testing.T) {
	SetupForTest()
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
	SetupForTest()
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
	SetupForTest()
	pwd, err := GetPwdByUid(1)
	fmt.Println(pwd, err)
	if err != nil {
		t.Fatal(err.Error())
	}
}
