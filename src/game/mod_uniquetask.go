package game

/*唯一任务模块 该模块在一个账号的生命周期内只能完成一次*/

type TaskInfo struct {
	TaskId int
	State  int // 任务状态
}

type ModUniqueTask struct {
	MyTaskInfo map[int]*TaskInfo
}

// IsTaskCompleted 任务是否完成
func (mt *ModUniqueTask) IsTaskCompleted(taskId int) bool {
	task, ok := mt.MyTaskInfo[taskId]
	if !ok {
		return false
	}
	return task.State == TASK_STATE_FINISH
}
