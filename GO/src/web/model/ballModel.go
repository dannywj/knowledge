package model

import (
	"fmt"
	db "web/database"
)

func ResetUserBall(guid string) bool {
	for i := 1; i <= 4; i++ {
		// 设置lasttime
		key := fmt.Sprintf("planting:user:ball:lasttime:%v:%v", guid, i)
		key_new := fmt.Sprintf("planting:user:ball:lasttime_new:%v:%v", guid, i)
		db.GlobalRedisClient.Del(key, key_new)

		// 清空被偷数
		key_stolen := fmt.Sprintf("planting:user:ball:stolen:energy:%v:%v", guid, i)
		db.GlobalRedisClient.Set(key_stolen, 0, 0)
		// 清空偷取记录
		key_collect := fmt.Sprintf("planting:user:ball:collect:%v:%v", guid, i)
		db.GlobalRedisClient.Del(key_collect)
		fmt.Println(key, key_new, key_stolen, key_collect)
	}
	return true
}
