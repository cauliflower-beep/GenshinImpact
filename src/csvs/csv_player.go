package csvs

import (
	"genshin-impact/src/utils"
)

// ConfigPlayerLevel 任务等级配置
type ConfigPlayerLevel struct {
	PlayerLevel int `json:"PlayerLevel"`
	PlayerExp   int `json:"PlayerExp"`
	WorldLevel  int `json:"WorldLevel"`
	ChapterId   int `json:"ChapterId"`
}

var (
	ConfigPlayerLevelSlice []*ConfigPlayerLevel
)

func init() {
	// 这里的路径要从项目根目录开始，因为是在init中处理的
	// 不知道路径的话，os打印一下当前路径就ok
	utils.GetCsvUtilMgr().LoadCsv("src/csvs/csv/PlayerLevel", &ConfigPlayerLevelSlice)
	return
}

// GetCurrentLevelConfig 获取当前等级的配置
func GetCurrentLevelConfig(level int) *ConfigPlayerLevel {
	if level < 0 || level >= len(ConfigPlayerLevelSlice) {
		return nil
	}
	return ConfigPlayerLevelSlice[level-1]
}
