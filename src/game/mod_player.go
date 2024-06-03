package game

import (
	"fmt"
	"genshin-impact/src/csvs"
)

type ModPlayer struct {
	Uid          int64       // uid
	Icon         int         // 头像
	IdCard       int         // 名片
	Signature    string      // 签名
	NickName     string      // 昵称
	PlayerLevel  int         // 冒险等级
	PlayerExp    int         // 冒险经验
	WorldLevel   int         // 世界等级
	WorldLevelCd int64       // 冷却时间  可以临时设置世界等级，冷却时间到了之后可以恢复
	Birthday     int         // 生日
	ShowTeam     map[int]int // 展示阵容 英雄id：英雄等级
	ShowCard     []int       // 展示名片
	// 看不见的字段 远比看得见的字段多
	IsProhibit int // 是否黑名单 根据这个值决定后续是否获取其他模块的数据 不用布尔值是方便扩展
	IsGm       int //该玩家是否gm号 游戏测试阶段还是需要大量的gm功能去辅助的
}

// setIcon 设置头像
func (mp *ModPlayer) setIcon(iconId int, player *Player) bool {

	// 这里必须做校验，试想，假设客户端被破解了，添加了一个非法的icon，请求打过来，如果服务器不做验证，肯定是存在风险的
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

// setIdcard 设置名片
func (mp *ModPlayer) setIdcard(idcard int, player *Player) bool {

	if !player.ModIdcard.isHasIdcard(idcard) {
		fmt.Printf("can not find %d !", idcard)
		return false
	}
	player.ModPlayer.IdCard = idcard
	fmt.Println("current idcard:", idcard)
	return true
}

// setName 设置昵称
func (mp *ModPlayer) setName(name string, player *Player) {
	// 违禁词处理
	if GetMgrBanWord().IsBanWord(name) {
		fmt.Println("set name failed. name contains ban words")
		return
	}

	player.ModPlayer.NickName = name
	fmt.Println("[current name]|", player.ModPlayer.NickName)
}

// setSignature 设置签名
func (mp *ModPlayer) setSignature(signature string, player *Player) {
	// 违禁词处理
	if GetMgrBanWord().IsBanWord(signature) {
		fmt.Println("set signature failed. signature contains ban words")
		return
	}

	player.ModPlayer.Signature = signature
	fmt.Println("[current signature]|", player.ModPlayer.Signature)
}

// AddExp 增加经验
func (mp *ModPlayer) AddExp(exp int, p *Player) {
	// 这是一个内置接口，不需要做经验值校验
	mp.PlayerExp += exp

	// 经验是可以一直累加的，直到做了某个突破任务之后，一口气加上去，所以有可能连升多级
	for {
		config := csvs.GetCurrentLevelConfig(mp.PlayerLevel)
		if config == nil {
			// 日志打出来让策划去查配置是不是有问题
			break
		}
		if config.PlayerExp == 0 {
			break
		}

		// 如果存在突破任务并且没有完成，直接跳出
		if config.ChapterId > 0 && !p.ModUniqueTask.IsTaskCompleted(config.ChapterId) {
			break
		}

		if mp.PlayerExp >= config.PlayerExp {
			mp.PlayerLevel++
			mp.PlayerExp -= config.PlayerExp
		} else {
			break
		}
	}
	fmt.Printf("[当前等级]^%d|[当前经验]^%d\n", mp.PlayerLevel, mp.PlayerExp)
}
