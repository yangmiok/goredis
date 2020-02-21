package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "13.52.237.35:6379")
	defer c.Close()
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	//自定义操作
	//如果 key 已经持有其他值， SET 就覆写旧值， 无视类型。
	//res, err := c.Do("set", "name", "taylor")
	//将键 key 的值设置为 value ， 并将键 key 的生存时间设置为 seconds 秒钟
	//res, err := c.Do("SETEX", "cache_user_id", "60", "10086")
	//如果键 key 已经存在并且它的值是一个字符串， APPEND 命令将把 value 追加到键 key 现有值的末尾。
	//res, err := c.Do("APPEND", "name", "-zhu")
	res, err := c.Do("INCR", "total")
	if err != nil {
		panic(err)
	}
	fmt.Print(res)
	//username, err := redis.String(c.Do("GET", "name"))
	//if err != nil {
	//	fmt.Println("redis get failed:", err)
	//} else {
	//	fmt.Printf("Get name: %v \n", username)
	//}

}
