package game

const (
	// 原神在做任务的时候是没办法联机的，所以这里定义了几个任务状态
	TASK_STATE_INIT   = iota
	TASK_STATE_DOING  // 玩家进入副本
	TASK_STATE_FINISH // 玩家完成副本
)

type Player struct {
	ModPlayer     *ModPlayer     // 基础模块
	ModIcon       *ModIcon       // 头像模块
	ModIdcard     *ModIdcard     // 名片模块
	ModUniqueTask *ModUniqueTask // 任务模块
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
	player.ModUniqueTask = new(ModUniqueTask)

	/******************mod init end****************************************/
	player.ModPlayer.PlayerLevel = 1 // 玩家等级初始为1级，否则后面从配置里面取等级配置是取不到的
	return player
}

// RecvSetIcon 设置头像
func (p *Player) RecvSetIcon(iconId int) {
	p.ModPlayer.setIcon(iconId, p)
}

// RecvSetIdcard 设置名片
func (p *Player) RecvSetIdcard(idcard int) {
	p.ModPlayer.setIdcard(idcard, p)
}

// RecvSetName 设置玩家名字
func (p *Player) RecvSetName(name string) {

	p.ModPlayer.setName(name, p)
}

// RecvSetSignature 设置签名
func (p *Player) RecvSetSignature(signature string) {

	p.ModPlayer.setSignature(signature, p)
}
