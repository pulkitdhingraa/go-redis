package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func must(err error){
	if err != nil {
		panic(err)
	}
}

func json(rdb *redis.Client, ctx context.Context) {
	val, err := rdb.JSONSet(ctx, "sector", "$", "\"IT\"").Result()
	must(err)
	fmt.Println(val)

	val2, err := rdb.JSONGet(ctx, "sector", "$").Result()
	must(err)
	fmt.Println(val2)

	val3, err := rdb.JSONStrLen(ctx, "sector", "$").Result()
	must(err)
	fmt.Println(*val3[0])

	val4, err := rdb.JSONStrAppend(ctx, "sector", "$", "\" - Information Technology\"").Result()
	must(err)
	fmt.Println(*val4[0])

	val5, err := rdb.JSONGet(ctx, "sector", "$").Result()
	must(err)
	fmt.Println(val5)
}

func jsonArray(rdb *redis.Client, ctx context.Context) {
	val, err := rdb.JSONSet(ctx, "user", "$", []interface{}{
			"pd",
			map[string]interface{}{"score": 0},
			nil,
		},
	).Result()
	must(err)
	fmt.Println(val)

	// val2, err := rdb.JSONGet(ctx, "user", "$").Result()
	// val2, err := rdb.JSONGet(ctx, "user", "$[0]").Result()
	// val2, err := rdb.JSONGet(ctx, "user", "$[1].score").Result()
	val2, err := rdb.JSONDel(ctx, "user", "$.[-1]").Result()

	must(err)
	fmt.Println(val2)
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

	// json(rdb, ctx)
	jsonArray(rdb, ctx)
}