package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
)

func main() {
	//normal()
	//args()
	//fbool()
	//sortset()
	sortset1()
}

func normal() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	resp, err := conn.Do("SET", "myKey", "abcd")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	resp, err = conn.Do("GET", "myKey")
	if err != nil {
		fmt.Println(err)
	}
	value, _ := redis.String(resp, err)
	fmt.Println(value)
}

func args() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var p1, p2 struct {
		Title string `redis:"title"`
		Author string `redis:"author"`
		Body string `redis:"body"`
	}
	p1.Title = "Example"
	p1.Author = "Gary"
	p1.Body = "Hello"
	
	if _, err := conn.Do("HMSET", redis.Args{}.Add("id1").AddFlat(&p1)...); err != nil {

		fmt.Println(err)
		return
	}

	m := map[string]string {
		"title": "Example2",
		"author": "Steve",
		"body": "Map",
	}
	if _, err := conn.Do("HMSET", redis.Args{}.Add("id2").AddFlat(m)...); err != nil {
		fmt.Println(err)
		return
	}

	for _, id := range []string{"id1", "id2"} {
		v, err := redis.Values(conn.Do("HGETALL", id))
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := redis.ScanStruct(v, &p2); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v\n", p2)
	}
}

func fbool() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")	
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	c.Do("SET", "foo", 1)
	exists, _ := redis.Bool(c.Do("EXISTS", "foo"))
	fmt.Printf("%#v\n", exists)
}

func sortset() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")	
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	for i, member := range []string{"red", "blue", "green"} {
		c.Do("ZADD", "zset", i, member)
	}	
	c.Do("ZADD", "zset", 1, "uio")
	resp, err := c.Do("ZCARD", "zset")
	if err != nil {
		log.Fatal(err)
	}
	count, err := redis.Int(resp, err)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)

	resp, err = c.Do("ZCOUNT", "zset", 1, 1)
	count, err = redis.Int(resp, err)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
func sortset1() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	//c.Do("DEL", "myset")
	// temp := book {
	// 	Title: "uio",
	// 	Author: "uuu",
	// 	Body: "qqqqqqqqqqqqq",
	// }
	// if _, err := c.Do("ZADD", "myset", 0, redis.Args{}.AddFlat(&temp)); err != nil {
	// 	log.Fatal(err)
	// }
	// // resp, err := c.Do("ZCARD", "myset")
	// // count, err := redis.Int(resp, err)
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// // fmt.Println(count)
	// resp, err := c.Do("ZRANGE", "myset", 0, -1)
	// values, err := redis.Values(resp, err)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// var bs []book
	// if err := redis.ScanStruct(values, &bs); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(bs)
}

type book struct {
	Title string
	Author string
	Body string
}


