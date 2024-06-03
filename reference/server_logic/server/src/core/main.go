package main

import (
	"server/csvs"
	"server/game"
)

func main() {
	// https://www.bilibili.com/
	// B站： golang大海葵

	// 专题：从0开始实现原神的服务器业务逻辑
	// 当前模块：背包    增加角色模块
	// 1  物品识别
	// 2  物品增加
	// 3  物品消耗
	// 4  物品使用
	// 5  角色模块--》头像模块

	//**********************************************************
	// 加载配置
	csvs.CheckLoadCsv()
	go game.GetManageBanWord().Run()

	//fmt.Printf("数据测试----start\n")
	playerTest := game.NewTestPlayer()
	go playerTest.Run()

	for{

	}

	//ticker := time.NewTicker(time.Second * 10)
	//for {
	//	select {
	//	case <-ticker.C:
	//		playerTest := game.NewTestPlayer()
	//		go playerTest.Run()
	//	}
	//}

	return
}

func playerLoadConfig(player *game.Player) {
	for i := 0; i < 1000000; i++ {
		config := csvs.ConfigUniqueTaskMap[10001]
		if config != nil {
			println(config.TaskId)
		}
	}
}
