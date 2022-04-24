package game

type Player struct {
	ModPlayer *ModPlayer
	ModIcon   *ModIcon
	ModIdcard *ModIdcard
}

func NewTestPlayer() *Player {
	/******************mod init begin**************************************/
	/*
		这个地方必须要初始化，否则就是一个空指针；那当前玩家这个线程就会崩掉
		服务器本身是不会崩掉的，因为golang多线程的话，只是玩家这个线程崩了，客户端会显示断线重连；其他玩家还是正常的
	*/
	player := new(Player)
	player.ModPlayer = new(ModPlayer)
	player.ModIcon = new(ModIcon)
	player.ModIdcard = new(ModIdcard)

	/******************mod init end****************************************/
	player.ModPlayer.Icon = 0
	player.ModPlayer.IdCard = 0
	return player
}

// RecvSetIcon
/*
类似设置头像这种请求，一定是外部可见的，因为他要跟客户端交互.
还有一些接口是不能暴露的；例如玩家等级的提升，一定是通过服务器内部的逻辑去处理的，不能说客户端请求过来，想设置几级就设置几级的
换言之，对外接口就是与客户端交互的接口；对内接口就是服务器内部逻辑调用的接口
*/
// 对外接口
func (self *Player) RecvSetIcon(iconId int) {
	/*
		客户端直接跟Player玩家主体打交道.
		Player收到消息之后再去调用其他模块来实现功能
		这样写的好处是，当你想要调试的时候，第一个断点肯定是打在这个地方，首先确保收到的是想要的数据；进而再去模块内部检查
	*/
	self.ModPlayer.setIcon(iconId, self)
}

func (self *Player) RecvSetIdcard(idcard int) {
	self.ModPlayer.setIdcard(idcard, self)
}
