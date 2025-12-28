package main

import (
	"context"
	"fmt"
	"ostadbun/adaptor/redisAdaptor"
	"ostadbun/pkg/hash"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

func main() {
	t := time.Now()
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}

	rdb := redisAdaptor.New()

	email := hash.Hasher("mohammadmosayyebnezhad@gmail.com")

	fmt.Println(email)
	code := "0951680d-05ce-427b-96a2-70a32bef0f55"

	key := fmt.Sprintf("osbn-u-s:%s:%s", email, code)

	//type sw struct {
	//	Name   string
	//	Os     string
	//	Client string
	//}
	//
	//s := sw{
	//	Name:   "ostadbun",
	//	Os:     "linux",
	//	Client: "ostadbun",
	//}
	//
	//v, _ := json.Marshal(s)
	//
	//err := rdb.Set(ctx, key, v, 0).Err()
	//if err != nil {
	//	log.Fatal(err)
	//}

	members, err := rdb.Keys(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}

	for _, member := range members {
		fmt.Println(member)

		str, err := rdb.Get(ctx, member).Result()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(str)
	}

	fmt.Println("Members:", members, time.Since(t))

}
