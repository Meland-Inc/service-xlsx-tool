package xlsx_export

import (
	"encoding/json"
	"fmt"
	"time"

	"game-message-core/proto"
	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"gorm.io/gorm"
)

// Task.xlsx

var taskTableRows = make(map[int32]xlsxTable.TaskTableRow)

func intArrToTaskRowOption(optType proto.TaskOptionType, arr []int, isChance bool) (*xlsxTable.TaskRowOption, error) {
	if len(arr) < 1 || (isChance && len(arr) < 2) {
		return nil, fmt.Errorf("option invalid")
	}

	opt := &xlsxTable.TaskRowOption{OptionType: optType}
	if isChance {
		opt.Chance = int32(arr[len(arr)-1])
	}

	switch optType {
	case proto.TaskOptionType_HandInItem:
		// 上交若干数量的指定道具(寻物任务) "1010001,1;1010002,1;1010003,1;1010004,1",
		opt.Param1 = int32(arr[0])
		opt.Param2 = int32(arr[1])

	case proto.TaskOptionType_UseItem:
		// 使用若干数量的指定道具 "1010001,1;1010002,1;1010003,1;1010004,1", // id,num;id,num;...
		opt.Param1 = int32(arr[0])
		opt.Param2 = int32(arr[1])

	case proto.TaskOptionType_PickUpItem:
		// 获得若干数量的指定道具 "1010001,1;1010002,1;1010003,1;1010004,1", // id,num;id,num;...
		opt.Param1 = int32(arr[0])
		opt.Param2 = int32(arr[1])

	case proto.TaskOptionType_KillMonster:
		// 杀死若干数量的指定怪物 "9900002,1;9900003,1;9900004,1", // cid, num; cid,num;...
		opt.Param1 = int32(arr[0])
		opt.Param2 = int32(arr[1])

	case proto.TaskOptionType_UserLevel:
		// 角色达到某等级"10", // level;
		opt.Param1 = int32(arr[0])

	case proto.TaskOptionType_TargetSlotLevel:
		// 指定插槽达到某等级 "1,4;5,3", // 插槽位置序号,等级;插槽位置序号,等级;...
		opt.Param1 = int32(arr[0])
		opt.Param2 = int32(arr[1])

	case proto.TaskOptionType_SlotLevelCount:
		// 指定数量插槽都达到某等级 "3,4", // 等级,插槽数量
		opt.Param1 = int32(arr[0])
		opt.Param2 = int32(arr[1])

	case proto.TaskOptionType_CraftSkillLevel:
		// 指定合成技能达到某等级 "20001,3;20002,2", //合成技能id,等级;合成技能id,等级;....
		opt.Param1 = int32(arr[0])
		opt.Param2 = int32(arr[1])

	case proto.TaskOptionType_UseRecipe:
		// 使用指定的配方合成 "20001,3;20002,2", //配方id,次数;配方id,次数;.....
		opt.Param1 = int32(arr[0])
		opt.Param2 = int32(arr[1])

	case proto.TaskOptionType_RecipeUseCount:
		// 累计合成多少次 "10", //次数
		opt.Param1 = int32(arr[0])

	case proto.TaskOptionType_TaskListTypeCount:
		// 完成若干数量的指定类型任务链 "1,3;2,2", //任务链类型,数量;任务链类型,数量;.....
		opt.Param1 = int32(arr[0])
		opt.Param2 = int32(arr[1])

	case proto.TaskOptionType_TargetPosition:
		// 到达指定坐标点指定半径范围内的区域 "111,3,111,2;222,3,222;333,3,333", // x,y,z,半径;
		opt.Param1 = int32(arr[0])
		opt.Param2 = int32(arr[1])
		opt.Param3 = int32(arr[2])
		opt.Param4 = int32(arr[3])
	}

	return opt, nil
}

func ParseTaskOptionsJson(id int32, js string, isChance bool) (optionList *xlsxTable.TaskRowOptionList, err error) {
	xlsxOptions := &xlsxTable.TaskXlsxOptions{}
	err = json.Unmarshal([]byte(js), xlsxOptions)
	if err != nil {
		return nil, err
	}

	optionList = &xlsxTable.TaskRowOptionList{}
	for _, option := range xlsxOptions.Options {
		intArrArr := excel.ParseIntSliceSliceValue(option.Value)
		for _, intArr := range intArrArr {
			opt, err := intArrToTaskRowOption(option.OptionType, intArr, isChance)
			if err != nil {
				serviceLog.Error(err.Error())
				continue
			}
			optionList.Options = append(optionList.Options, *opt)
			optionList.ChanceSum += opt.Chance
		}
	}

	return optionList, nil
}

func ParseTask(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}
		setting := xlsxTable.TaskTableRow{
			Id:         excel.IntToInt32(row["id"]),
			Level:      excel.IntToInt32(row["level"]),
			Name:       excel.StringToString(row["name"]),
			SubSystem:  excel.IntSliceToJsonStr(row["subSystem"]),
			RewardId:   excel.IntToInt32(row["itemReward"]),
			RewardExp:  excel.IntToInt32(row["expReward"]),
			Difficulty: excel.IntToInt32(row["difficulty"]),
		}

		// 解析designateOptions
		designateOptionsJs := excel.StringToString(row["designateOptions"])
		desOptList, err := ParseTaskOptionsJson(setting.Id, designateOptionsJs, false)
		if err != nil {
			err = fmt.Errorf(" Task.xlsx id[%v] designateOptions 配置错误", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}
		if err = setting.SetDesignateOptions(desOptList); err != nil {
			serviceLog.Error(err.Error())
			continue
		}

		// 解析designateOptions
		chanceOptionsJs := excel.StringToString(row["chanceOptions"])
		chaOptList, err := ParseTaskOptionsJson(setting.Id, chanceOptionsJs, true)
		if err != nil {
			err = fmt.Errorf(" Task.xlsx id[%v] chanceOptions 配置错误", setting.Id)
			serviceLog.Error(err.Error())
			continue
		}
		if err = setting.SetChanceOptions(chaOptList); err != nil {
			serviceLog.Error(err.Error())
			continue
		}

		if err = setting.Check(); err != nil {
			serviceLog.Error(err.Error())
			continue
		}

		taskTableRows[setting.Id] = setting
	}

	return err
}

func TaskSaveToDB(db *gorm.DB, curSecUtc time.Time) {
	list := []xlsxTable.TaskTableRow{}
	for _, Task := range taskTableRows {
		Task.CreatedAt = curSecUtc
		list = append(list, Task)
	}

	WriterToDB(db, curSecUtc, &xlsxTable.TaskTableRow{}, len(list), list)
}
