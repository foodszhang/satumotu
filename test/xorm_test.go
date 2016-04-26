package xorm_test

import (
	"satumatu"
	"testing"
	"time"
)

type User struct {
	Id            int64
	Username      string `xorm:"notnull unique"`
	Password      string
	DateCreated   time.Time `xorm:"created"`
	DateLastLogin time.Time
}

func TestCreate(t *testing.T) {
	err := satumatu.DBInit("satumatu", "satumatu", "ssss", "", "")

	if err != nil {
		t.Fatal(err)
	}
	db := satumatu.DB
	err = db.Sync2(new(User))
	user := new(User)
	has, err := db.Where("username = ?", "foods").Get(user)
	if err != nil {
		t.Fatal(err)
	}
	if !has {
		user.Username = "foods"
		user.Password = "ffff"
		affect, err := db.Insert(user)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(affect)

	}
	nuser := new(User)
	has, err = db.Where("username = ?", "foods").Get(nuser)
	t.Log(has, nuser)
}
