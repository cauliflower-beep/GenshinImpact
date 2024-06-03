package main

import (
	"genshin-impact/src/csvs"
	_ "genshin-impact/src/csvs"
	"genshin-impact/src/game"
)

func main() {

	//加载配置
	csvs.CheckLoadCsv()

	//player := game.NewTestPlayer()
	//fmt.Println(player.ModIdcard)
	//player.RecvSetIcon(1)   // 胡桃icon
	//player.RecvSetIcon(3)   // 钟离icon
	//player.RecvSetIdcard(2) // 枫叶

	//设置名字和签名
	//player.RecvSetName("小新")
	//player.RecvSetName("风间")
	//player.RecvSetName("小新开挂")
	//player.RecvSetName("小新原神")
	//
	//player.RecvSetSignature("大象，大象你的鼻子怎么那么长")
	//player.RecvSetSignature("我爱原神")
	//player.RecvSetSignature("我卢本伟没有开挂！")
	//go game.GetMgrBanWord().Run()

	playerGM := game.NewTestPlayer()

	playerGM.ModPlayer.AddExp(10000000, playerGM) // 由于存在突破任务，即使直接给10000000经验，也会卡在第一个突破任务的等级
	//tricker := time.NewTicker(time.Second)
	//for {
	//	select {
	//	case <-tricker.C:
	//		if time.Now().Unix()%3 == 0 {
	//			playerGM.ModPlayer.AddExp(5000, playerGM)
	//		}
	//	}
	//}

}
