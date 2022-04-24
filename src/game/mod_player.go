package game

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
