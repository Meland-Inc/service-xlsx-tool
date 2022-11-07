package xlsx_make

import (
	"game-message-core/proto"
	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/xlsx_make/data"
)

/*
	每种任务项配置示例， 下面的示例为 固定的任务项， 权重任务项需要在下面的value项中的 ";" 前面添加 ,权重值;

	TaskOptionType_HandInItem   			// 上交若干数量的指定道具(寻物任务)
	{
		OptionType: proto.TaskOptionType_HandInItem,
		Value:      "1010001,1;1010002,1;1010003,1;1010004,1", // id,num;id,num;...
	},

	TaskOptionType_UseItem 					// 使用若干数量的指定道具
	{
		OptionType: proto.TaskOptionType_UseItem,
		Value:      "1010001,1;1010002,1;1010003,1;1010004,1", // id,num;id,num;...
	},

	TaskOptionType_PickUpItem   			// 获得若干数量的指定道具
	{
		OptionType: proto.TaskOptionType_PickUpItem,
		Value:      "1010001,1;1010002,1;1010003,1;1010004,1", // id,num;id,num;...
	},

	TaskOptionType_KillMonster				// 杀死若干数量的指定怪物
	{
		OptionType: proto.TaskOptionType_KillMonster,
		Value:      "9900002,1;9900003,1;9900004,1", // cid, num; cid,num;...
	},

	TaskOptionType_UserLevel 				// 角色达到某等级
	{
		OptionType: proto.TaskOptionType_UserLevel,
		Value:      "10", // level;
	},

	TaskOptionType_TargetSlotLevel 			// 指定插槽达到某等级
	{
		OptionType: proto.TaskOptionType_TargetSlotLevel,
		Value:      "1,4;5,3", // 插槽位置序号,等级;插槽位置序号,等级;...
	},

	TaskOptionType_SlotLevelCount			// 指定数量插槽都达到某等级
	{
		OptionType: proto.TaskOptionType_SlotLevelCount,
		Value:      "3,4", // 等级,插槽数量
	},

	TaskOptionType_CraftSkillLevel			// 指定合成技能达到某等级
	{
		OptionType: proto.TaskOptionType_CraftSkillLevel,
		Value:      "20001,3;20002,2", //合成技能id,等级;合成技能id,等级;....
	},

	TaskOptionType_RecipeUse				// 使用指定的配方合成
	{
		OptionType: proto.TaskOptionType_RecipeUse,
		Value:      "20001,3;20002,2", //配方id,次数;配方id,次数;.....
	},

	TaskOptionType_TaskListTypeCount	   // 完成若干数量的指定类型任务链
	{
		OptionType: proto.TaskOptionType_TaskTypeCount,
		Value:      "1,3;2,2", //任务链类型,数量;任务链类型,数量;.....
	},

	TaskOptionType_TargetPosition			// 到达指定坐标点指定半径范围内的区域
	{
		OptionType: proto.TaskOptionType_TargetPosition,
		Value:      "111,3,111,2;222,3,222;333,3,333", // x,y,z,半径;
	},
*/

func MakeTaskXlsxData() {
	// -------------------- task 1000X --------------------------------------------
	MakeTask1000N()
	// -------------------- task 2000X --------------------------------------------
	// MakeTask2000N()
	// -------------------- task 3000X --------------------------------------------
	// MakeTask3000N()
}

func MakeTask1000N() {
	idBegin, idEnd := int32(10001), int32(10003)

	// 固定的任务项
	designateOptions := make(map[int32][]xlsxTable.TaskXlsxRowOption)
	designateOptions[10001] = []xlsxTable.TaskXlsxRowOption{
		{
			OptionType: proto.TaskOptionType_TargetPosition,
			Value:      "111,111,3000;222,222,3000;333,333,4000",
		},
	}
	designateOptions[10002] = designateOptions[10001]
	designateOptions[10003] = designateOptions[10001]

	// 按权重随机的任务项
	chanceOptions := make(map[int32][]xlsxTable.TaskXlsxRowOption)
	chanceOptions[10001] = []xlsxTable.TaskXlsxRowOption{
		{
			OptionType: proto.TaskOptionType_KillMonster,
			Value:      "9900002,1,200;9900003,1,200;9900004,1,600", // cid, num,权重; cid,num，权重;...
		},
	}
	chanceOptions[10002] = chanceOptions[10001]
	chanceOptions[10003] = chanceOptions[10001]

	for id := idBegin; id <= idEnd; id++ {
		row := data.TaskXlsxRow{
			Id:               id,
			Level:            0,
			Name:             "Traveling",
			Decs:             "Go to the %% position",
			Details:          "For the sake of love and peace, please go to the designated place",
			SubSystem:        "1",
			ItemReward:       0,
			ExpReward:        100,        //递增 n + id % 10000 * 增量
			Difficulty:       id % 10000, // id % 10000 + 1
			DesignateOptions: designateOptions[id],
			ChanceOptions:    chanceOptions[id],
		}

		// // 指定修改某个id 的某个数据
		// if id == 10002 {
		// 	row.ExpReward = 99
		// 	row.Level = 1
		// }

		AddTaskXlsxData(row)
	}
}
