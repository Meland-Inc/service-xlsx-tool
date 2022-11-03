package data

import (
	"encoding/json"
	"fmt"
	xlsxTable "game-message-core/xlsxTableData"
)

var TaskXlsxFileHeader = [][]string{
	{"任务ID", "任务等级", "任务名", "任务描述", "任务详情", "任务链体系", "任务参数json", "任务道具奖励Id", "任务Exp奖励", "难度系数（预留）", "激活任务"},
	{"id", "level", "name", "decs", "details", "subSystem", "parameter", "itemReward", "expReward", "difficulty", "nextTaskId"},
	{"int", "int", "string", "string", "string", "int[]", "string", "int", "int", "int", "int"},
}

type TaskXlsxRow struct {
	Id         int32                         `json:"id"`
	Level      int32                         `json:"level"`
	Name       string                        `json:"name"`
	Decs       string                        `json:"decs"`
	Details    string                        `json:"details"`
	SubSystem  string                        `json:"subSystem"`
	ItemReward int32                         `json:"itemReward"`
	ExpReward  int32                         `json:"expReward"`
	Difficulty int32                         `json:"difficulty"`
	NextTaskId int32                         `json:"nextTaskId"`
	Options    []xlsxTable.TaskXlsxRowOption `json:"-"`
}

func (p *TaskXlsxRow) GetParameter() (string, error) {
	if len(p.Options) < 1 {
		return "", fmt.Errorf("task options is empty")
	}

	parameter := xlsxTable.TaskXlsxOptions{Options: p.Options}
	bs, err := json.Marshal(parameter)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

func (p *TaskXlsxRow) ToStringList() ([]string, error) {
	parameter, err := p.GetParameter()
	if err != nil {
		return nil, err
	}

	// 数据顺序一定不能错
	//{"id", "level", "name", "decs", "details", "subSystem", "parameter", "itemReward", "expReward", "difficulty", "nextTaskId"},
	strSlice := []string{
		fmt.Sprint(p.Id),
		fmt.Sprint(p.Level),
		p.Name,
		p.Decs,
		p.Details,
		p.SubSystem,
		parameter,
		fmt.Sprint(p.ItemReward),
		fmt.Sprint(p.ExpReward),
		fmt.Sprint(p.Difficulty),
		fmt.Sprint(p.NextTaskId),
	}
	return strSlice, nil
}
