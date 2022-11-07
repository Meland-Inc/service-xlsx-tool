package data

import (
	"encoding/json"
	"errors"
	"fmt"
	xlsxTable "game-message-core/xlsxTableData"
)

var TaskXlsxFileHeader = [][]string{
	{"任务ID", "任务等级", "任务名", "任务描述", "任务详情", "任务链体系", "指定任务选项", "权重任务选项", "任务道具奖励Id", "任务Exp奖励", "难度系数（预留）"},
	{"id", "level", "name", "decs", "details", "subSystem", "designateOptions", "chanceOptions", "itemReward", "expReward", "difficulty"},
	{"int", "int", "string", "string", "string", "int[]", "string", "string", "int", "int", "int"},
}

type TaskXlsxRow struct {
	Id               int32                      `json:"id"`
	Level            int32                      `json:"level"`
	Name             string                     `json:"name"`
	Decs             string                     `json:"decs"`
	Details          string                     `json:"details"`
	SubSystem        string                     `json:"subSystem"`
	ItemReward       int32                      `json:"itemReward"`
	ExpReward        int32                      `json:"expReward"`
	Difficulty       int32                      `json:"difficulty"`
	DesignateOptions []xlsxTable.TaskXlsxOption `json:"-"`
	ChanceOptions    []xlsxTable.TaskXlsxOption `json:"-"`
}

func (p *TaskXlsxRow) GetDesignateOptions() string {
	if len(p.DesignateOptions) < 1 {
		return ""
	}

	opts := xlsxTable.TaskXlsxOptions{Options: p.DesignateOptions}
	bs, err := json.Marshal(opts)
	if err != nil {
		panic(err)
	}

	return string(bs)
}

func (p *TaskXlsxRow) GetChanceOptions() string {
	if len(p.ChanceOptions) < 1 {
		return ""
	}

	opts := xlsxTable.TaskXlsxOptions{Options: p.ChanceOptions}
	bs, err := json.Marshal(opts)
	if err != nil {
		panic(err)
	}

	return string(bs)
}

func (p *TaskXlsxRow) ToStringList() ([]string, error) {
	designateOptionsStr := p.GetDesignateOptions()
	chanceOptionsStr := p.GetChanceOptions()
	if designateOptionsStr == "" && chanceOptionsStr == "" {
		return nil, errors.New("task options is invalid")
	}

	// 数据顺序一定不能错
	//{"id", "level", "name", "decs", "details", "subSystem", "designateOptions", "chanceOptions", "itemReward", "expReward", "difficulty", "nextTaskId"},
	strSlice := []string{
		fmt.Sprint(p.Id),
		fmt.Sprint(p.Level),
		p.Name,
		p.Decs,
		p.Details,
		p.SubSystem,
		designateOptionsStr,
		chanceOptionsStr,
		fmt.Sprint(p.ItemReward),
		fmt.Sprint(p.ExpReward),
		fmt.Sprint(p.Difficulty),
	}
	return strSlice, nil
}
