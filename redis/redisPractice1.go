package main

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//连接redis
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis err: ", err)
		return
	}
	defer c.Close()

	// 对redis进行读写，写入的值永远不会过期
	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("read set failed: ", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis set failed: ", err)
	} else {
		fmt.Printf("Get mykey: %v\n", username)
	}

	//设置过期，可以使用SET的附加参数
	_, err = c.Do("SET", "mykey", "superWang", "EX", "5")
	if err != nil {
		fmt.Println("read set failed: ", err)
	}

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis set failed: ", err)
	} else {
		fmt.Printf("Get mykey: %v\n", username)
	}

	time.Sleep(8 * time.Second)
	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis set failed: ", err)
	} else {
		fmt.Printf("Get mykey: %v\n", username)
	}

}
