package xorm_test

import (
	"testing"
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type User struct {
	Id            int64
	Username      string `xorm:"notnull unique"`
	Password      string
	DateCreated   time.Time `xorm:"created"`
	DateLastLogin time.Time
}

func TestCreate(t *testing.T) {
	engine, err := xorm.NewEngine("postgres", "user=satumatu password=ssss dbname=satumatu sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	engine.SetMapper(core.GonicMapper{})
	err = engine.Sync2(new(User))
	user := new(User)
	has, err := engine.Where("username = ?", "foods").Get(user)
	if err != nil {
		t.Fatal(err)
	}
	if !has {
		affect, err := engine.Insert(user)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(affect)

	}
	nuser := new(User)
	has, err = engine.Where("username = ?", "foods").Get(nuser)
	t.Log(has, nuser)
}
