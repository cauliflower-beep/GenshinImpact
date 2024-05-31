package csvs

import "fmt"

// 其他的配置初始化都是有init()函数自动调用的
// 此外我们还要定义一个check函数来手动调用，表示初始化完成

// CheckLoadCsv 手动调用这个函数，则表示配置初始化完成
func CheckLoadCsv() {
	// init函数中不会做任何的逻辑处理 因为它的执行顺序与文件名有关 如果做逻辑处理还要跟文件名做关联，太牵强
	// 表的二次处理放在这里进行
	fmt.Println("load csv is done.")
}
