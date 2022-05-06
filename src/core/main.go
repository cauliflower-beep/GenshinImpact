package main

import (
	"fmt"
	_ "originalGod/src/csvs"
	"originalGod/src/game"
)

/*
先聊一聊设计理念的不同。
现在假设有200个玩家从客户端链接过来，
以8核 CPU 为例,C++的做法是分成1+7，一个作为模块管理，管理7个模块，那7个模块去处理200个玩家的请求，中间可能用到了消息队列；
但go的做法是直接起200个协程，因为协程是很轻量级的，一个只有几k，所以8核可以很轻易的处理大量的玩家协程
*/

/*
main函数里面要做的第一件事，往往是导入配置。
例如策划维护了一张excel的表格，我们服务器要做的事情就是利用导表工具，把这个excel转换成我们更好利用的格式，比如csv（更像一个数组），或者json（更像一个map）
*/
func main() {

	//加载配置
	player := game.NewTestPlayer()
	fmt.Println(player.ModIdcard)
	player.RecvSetIcon(1)   // 胡桃icon
	player.RecvSetIdcard(2) // 枫叶
}
