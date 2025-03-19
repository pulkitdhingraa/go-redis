package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func strings(rdb *redis.Client, ctx context.Context) {
	err := rdb.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)

	err = rdb.MSet(ctx, "a", "newtext", "b", "nextx").Err()
	if err != nil {
		panic(err)
	}

	val2, err := rdb.MGet(ctx, "a", "b").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val2)

	val3, err := rdb.LCS(ctx, &redis.LCSQuery{
		Key1: "a",
		Key2: "b",
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val3.MatchString)

	err = rdb.SetRange(ctx, "b", 4, "abc").Err()
	if err != nil {
		panic(err)
	}

	val4, _ := rdb.Get(ctx, "b").Result()
	fmt.Println(val4)
}

func main() {
	ctx := context.Background()

	// Connect to a redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
		Protocol: 2,
	})

	strings(rdb, ctx)
}

