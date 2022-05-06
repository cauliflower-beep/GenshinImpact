package game

/*
违禁词库应该分为两部分：
1.基础违禁词库；
2.扩充违禁词库，通过定时调用接口去扩充
*/
import (
	"fmt"
	"regexp"
	"time"
)

type MgrBanword struct {
	BanwordBase  []string //基础违禁词库
	BanwordExtra []string // 扩充违禁词库
}

func (self *MgrBanword) IsBanword(txt string) bool {

	for _, v := range self.BanwordBase {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			return match
		}
	}
	for _, v := range self.BanwordExtra {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			return match
		}
	}
	return false
}

// Run 定时扩充违禁词库，有个定时器，服务器启动的时候就会跟着启动
func (self *MgrBanword) Run() {
	ticker := time.NewTicker(time.Second * 1) // golang中的定时器本质上就是一个channel
	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now().Unix())
			if time.Now().Unix()%10 == 0 { //每10秒更新词库
				fmt.Println("update banword")
			} else {
				fmt.Println("standby...")
			}
		}
	}
}
