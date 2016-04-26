package auth

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"math/rand"
	"satumatu"
	"time"
)

// Token give prefix and expires for redis to store data
type Token struct {
	TokenPrefix  string
	RedisPrefix  string
	MaxExpiresIn int
}

func (token *Token) makeToken(salt string) string {
	//generate token and return
	t := time.Now().UnixNano()
	rand.Seed(t)
	random := rand.Int()
	origin := fmt.Sprintf("%s%d-%d-%s", token.TokenPrefix, t, random, salt)
	return origin
}

// Create change the
func (token *Token) Create(data interface{}, expires int) string {
	result, _ := json.Marshal(data)
	str := string(result)
	var tokenStr string
	var origin string
	for {
		origin = token.makeToken(str)

		//set key value ex time nx
		tokenStr = fmt.Sprintf("%x", sha1.Sum([]byte(origin)))
		reply, _ := satumatu.Redis.Do("SET", token.RedisPrefix+tokenStr, str, "ex", expires, "nx")
		if reply == "OK" {
			break
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}

	return tokenStr
}

// Delete data
func (*Token) Delete(token string) {
	satumatu.Redis.Do("DEL", token)
}

// Retrieve 出去之后再进行类型转化
func (token *Token) Retrieve(tokenStr string) []byte {
	result, _ := satumatu.Redis.Do("GET", token.RedisPrefix+tokenStr)
	dataBytes, b := result.([]byte)
	if !b {
		fmt.Println("转化失败")
	}
	return dataBytes
}
