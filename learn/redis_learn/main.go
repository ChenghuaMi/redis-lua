package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	//runLuaScript(rdb)
	//loadLuaScript(rdb)
	evalScript(rdb)
}

// 执行lua脚本
func runLuaScript(rdb *redis.Client) {
	luaScript := `
	local key = KEYS[1]
	local value = ARGV[1]
	redis.call('set',key,value)
	return redis.call('get',key)
`
	script := redis.NewScript(luaScript)
	result, err := script.Run(context.Background(), rdb, []string{"hello"}, "word").Result()
	fmt.Println("err:", err)
	fmt.Println("result:", result)
}

// 使用eval 执行 lua
func evalScript(rdb *redis.Client) {
	luaScript := `
	local key = KEYS[1]
	local value = ARGV[1]
	redis.call('SET',key,value)
	return redis.call('GET',key)
`
	result, err := rdb.Eval(context.Background(), luaScript, []string{"mchx"}, "michenghua....").Result()
	fmt.Println("result:", result, "err:", err)
}

// 使用 evalsha 执行lua
func loadLuaScript(rdb *redis.Client) error {
	luaScript := `
	local key = KEYS[1]
	local value = ARGV[1]
	redis.call('SET',key,value)
	return redis.call('GET',key)
`
	ctx := context.Background()
	sha, err := rdb.ScriptLoad(ctx, luaScript).Result()
	if err != nil {
		return err
	}

	//script := redis.NewScript(luaScript)

	result, err := rdb.EvalSha(ctx, sha, []string{"word"}, "golang").Result()

	fmt.Println("result:", result, "err:", err)
	return nil
}
