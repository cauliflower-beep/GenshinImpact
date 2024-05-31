package game

/*
违禁词库应该分为两部分：
1.基础违禁词库；
2.扩充违禁词库，通过定时调用接口去扩充
*/
import (
	"fmt"
	"genshin-impact/src/csvs"
	"regexp"
	"time"
)

// mgrBanWord 违禁词单例
var mgrBanWord *MgrBanWord

type MgrBanWord struct {
	BanWordBase  []string //基础违禁词库
	BanWordExtra []string // 扩充违禁词库 定时更新
}

func GetMgrBanWord() *MgrBanWord {
	// sync.Once 优化
	if mgrBanWord == nil {
		mgrBanWord = new(MgrBanWord)
	}
	return mgrBanWord
}

// IsBanWord 验证文本是否合法
func (mbw *MgrBanWord) IsBanWord(txt string) bool {

	for _, v := range mbw.BanWordBase {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			return match
		}
	}
	for _, v := range mbw.BanWordExtra {
		match, _ := regexp.MatchString(v, txt)
		fmt.Println(match, v)
		if match {
			return match
		}
	}
	return false
}

// Run 定时扩充违禁词库，有个定时器，服务器启动的时候就会跟着启动
func (mbw *MgrBanWord) Run() {
	mbw.BanWordBase = csvs.GetBanWordBase()
	ticker := time.NewTicker(time.Second * 1) // tricker核心是一个channel
	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now().Unix())
			if time.Now().Unix()%10 == 0 { //每10秒更新词库
				fmt.Println("update banWord")
			} else {
				fmt.Println("standby...")
			}
		}
	}
}
