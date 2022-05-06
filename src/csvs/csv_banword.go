package csvs

import "fmt"

/*
这里我们将决定在内存中，用什么样的配置去记录这些数据
*/

type ConfBanword struct {
	Id  int
	Txt string
}

var ConfBanwordSlice []*ConfBanword

//init 第一次调用这个包的时候进行初始化
func init() {
	/*
		一般来说，进公司就已经把导表的接口写好了，例如 loadCSV 之类的
		这里只展示纯业务逻辑的服务器，所以把违禁词写死，避免左边出现太多非go类的文件
	*/
	ConfBanwordSlice = append(ConfBanwordSlice,
		&ConfBanword{Id: 1, Txt: "外挂"},
		&ConfBanword{Id: 2, Txt: "原神"},
		&ConfBanword{Id: 3, Txt: "开卡"},
		&ConfBanword{Id: 4, Txt: "暴力"},
		&ConfBanword{Id: 5, Txt: "流血"},
	)
	fmt.Println("csv_banword init...")
}
