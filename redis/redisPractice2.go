// 批量写入读取

// MGET key [key …]
// MSET key value [key value …]

// 批量写入读取对象(Hashtable)
// HMSET key field value [field value …]
// HMGET key field [field …]

// 检测值是否存在
// EXISTS key

package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {

	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed: ", err)
	}

	//检测值是否存在
	is_key_exit, err := redis.Bool(c.Do("EXISTS", "mykey1"))
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Printf("exists or not: %v\n", is_key_exit)
	}

	//删除key
	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	_, err = c.Do("DEL", "mykey")
	if err != nil {
		fmt.Println("redis delete failed: ", err)
	}

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

}
