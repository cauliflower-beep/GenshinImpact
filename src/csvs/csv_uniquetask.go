package csvs

import "genshin-impact/src/utils"

type ConfigUniqueTask struct {
	TaskId    int `json:"TaskId"`    // 任务ID
	SortType  int `json:"SortType"`  // 任务分类 1-突破任务
	OpenLevel int `json:"OpenLevel"` // 开启等级，完成任务的时候需要验证一下
	TaskType  int `json:"TaskType"`  // 完成的条件
	Condition int `json:"Condition"`
}

var (
	ConfigUniqueTaskMap map[int]*ConfigUniqueTask
)

func init() {
	ConfigUniqueTaskMap = make(map[int]*ConfigUniqueTask)
	utils.GetCsvUtilMgr().LoadCsv("src/csvs/csv/UniqueTask", &ConfigUniqueTaskMap)
	return
}
