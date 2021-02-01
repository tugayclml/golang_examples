package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"time"
	"reflect"
)

var ctx = context.Background()

func main()  {

	fmt.Println("Go redis tutorial")

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	fmt.Println(reflect.TypeOf(client))

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	//setKeyValueToRedis("name", "tugay", 0, client)
	getValueWithKeyFromRedis("name", client)
	deleteKeyFromRedis("name", client)
	getValueWithKeyFromRedis("name", client)



}

func setKeyValueToRedis(key string, value string, expiration time.Duration, client *redis.Client)  {

	err := client.Set(key, value, expiration).Err()

	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println("Key saved.")
	}
}

func getValueWithKeyFromRedis(key string, client *redis.Client){

	val, err := client.Get(key).Result()
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(val)
	}
}

func deleteKeyFromRedis(key string, client *redis.Client)  {
	err := client.Del(key).Err()

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Key deleted")
	}
}

