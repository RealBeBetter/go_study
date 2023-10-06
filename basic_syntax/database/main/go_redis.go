/**
 * @author: Real
 * Date: 2022/11/15 0:44
 */
package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	// 首先需要安装需要的依赖
	// 安装之前，使用命令加速 go get
	// go env -w GO111MODULE=on
	// go env -w GOPROXY=https://goproxy.io,direct
	// 使用命令 go get github.com/garyburd/redigo/redis

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis connection failed..")
		return
	}

	// 写数据
	_, err = conn.Do("Set", "name", "Song Wei") // 参数列表：指令、键、值
	if err != nil {
		fmt.Println("redis set failed..")
		return
	}

	// 读数据
	r, err := redis.String(conn.Do("Get", "name"))
	fmt.Println("result = ", r)

	// Hash的操作
	_, err = conn.Do("HSet", "userhash01", "name", "Songwei") // 操作、哈希名、键、值
	_, err = conn.Do("HSet", "userhash01", "age", 23)

	r, err = redis.String(conn.Do("HGet", "userhash01", "name")) // 从哈希userhash01中读取name
	fmt.Println("hash name = ", r)
	age, err := redis.Int(conn.Do("HGet", "userhash01", "age")) // 读取int
	fmt.Println("hash age = ", age)

	// 一次写入或插入多个值
	_, err = conn.Do("MSet", "name", "SongWei", "home", "Hunan,Hengyang")
	multi_r, err := redis.Strings(conn.Do("MGet", "name", "home")) // 注意String多了个s，而且multi_r是[]string
	fmt.Println("Mset result is", multi_r)

	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("connection close failed..., err is", err)
		}
	}(conn)

	var pool *redis.Pool // 全局连接池指针
	pool = &redis.Pool{
		MaxIdle:     8,   // 最大空闲连接数
		MaxActive:   0,   // 最大连接数,0表示不限
		IdleTimeout: 100, // 最大空闲时间
		Dial: func() (redis.Conn, error) { // 产生连接的函数
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	conn = pool.Get() // 获取连接
	fmt.Println("Redis connection pool is", conn)

	defer func(pool *redis.Pool) {
		err := pool.Close()
		if err != nil {
			fmt.Println("connection close failed..., err is", err)
		}
	}(pool) // 连接池关闭
}
