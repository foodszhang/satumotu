package satumatu

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	// import pqdb
	_ "github.com/lib/pq"
)

// DB is the global engine to change databse
var DB *xorm.Engine

// Redis is the global engine to control redis
var Redis redis.Conn

// DBInit init the postgres database connect
func DBInit(dbname, username, password, host, port string) error {
	var err error
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	dbstring := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s  sslmode=disable", username,
		password, dbname, host, port)
	DB, err = xorm.NewEngine("postgres", dbstring)
	if err != nil {
		return err

	}
	DB.SetMapper(core.GonicMapper{})
	return nil
}

// RedisInit init redis connect
func RedisInit(host, port string, num int) error {
	var err error
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "6379"
	}
	Redis, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port), redis.DialDatabase(num))
	if err != nil {
		return err
	}
	return nil
}
