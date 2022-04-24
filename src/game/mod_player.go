package game

import "fmt"

/*
玩家基础模块信息，上线要发给玩家的：
1.uid
2.头像、名片
3.签名
4.名字
5.冒险等级、冒险阅历
6.世界等级、冷却时间(可以临时设置世界等级，冷却时间到了之后可以恢复)
7.生日
8.展示阵容、展示名片
*/

/*
这里的模块应该是对应数据库中的某一张表的;
玩家登录之后，就会从数据库中取出这些数据，发给客户端
*/
type ModPlayer struct {
	Uid          int64
	Icon         int
	IdCard       int
	Signature    string
	NickName     string
	PlayerLevel  int
	PlayerExp    int
	WorldLevel   int
	WorldLevelCd int64
	Birthday     int
	ShowTeam     map[int]int //英雄id：英雄等级
	ShowCard     []int
	// 看不见的字段
	IsProhibit int // 是否黑名单 根据这个值决定后续是否获取其他模块的数据
	IsGm       int //该玩家是否gm号 游戏测试阶段还是需要大量的gm功能去辅助的
}

/*
本项目中，接收客户端的消息一律用 Recv 开头.
原神的头像包，只展示所有拥有的英雄集合；原理上来说就是玩家登录之后，拿到服务器的数据，做一遍初始化就可以；
服务器方面，接收到客户端消息之后需要验证一下是否拥有该英雄；
设想一下，假设客户端被破解了，添加了一个非法的icon，请求打过来，如果服务器不做验证，肯定是存在风险的
*/

func (self *ModPlayer) setIcon(iconId int, player *Player) bool {

	if !player.ModIcon.isHasIcon(iconId) {
		/*
			没有请求的icon.
			公司内部应该是有一套机制的，来通知客户端操作非法
		*/
		fmt.Printf("can not find %d !", iconId)
		return false
	}
	player.ModPlayer.Icon = iconId
	fmt.Println("current icon:", iconId)
	return true
}

// setIdcard
/*
原神的名片展示跟头像不太一样，是把所有的名片都展示了出来，分为解锁和未解锁两部分；
这其实是策划的选择问题，如果是展示所有的名片，就是客户端读配置表，对名片做一个全局的初始化；
对于服务器来说并没有什么不同，都是发送拥有的物品
*/
func (self *ModPlayer) setIdcard(idcard int, player *Player) bool {

	if !player.ModIdcard.isHasIdcard(idcard) {
		fmt.Printf("can not find %d !", idcard)
		return false
	}
	player.ModPlayer.IdCard = idcard
	fmt.Println("current idcard:", idcard)
	return true
}
