package main

import (
	_ "genshin-impact/src/csvs"
	"genshin-impact/src/game"
)

func main() {

	//加载配置
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
	go game.GetMgrBanWord().Run()
	select {}

}
