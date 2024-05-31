package csvs

import "fmt"

/*
这里决定在内存中，用什么样的结构去处理这些配置数据
*/

type ConfBanword struct {
	Id  int
	Txt string
}

var ConfBanwordSlice []*ConfBanword

func init() {
	/*
		一般来说，进公司就已经把导表的接口写好了，例如 loadCSV 之类的，传数据结构进去，外加配置名，就能加载。
		这里只展示纯业务逻辑的服务器，所以把违禁词写死
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

func GetBanWordBase() []string {
	var banWords []string
	for _, banWord := range ConfBanwordSlice {
		banWords = append(banWords, banWord.Txt)
	}
	return banWords
}
