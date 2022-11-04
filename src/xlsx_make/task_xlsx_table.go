package xlsx_make

import (
	"game-message-core/proto"
	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/xlsx_make/data"
)

/*
	TaskOptionType_HandInItem   			// 上交若干数量的指定道具(寻物任务)
	TaskOptionType_UseItem 					// 使用若干数量的指定道具
	TaskOptionType_PickUpItem   			// 获得若干数量的指定道具
	TaskOptionType_KillMonster				// 杀死若干数量的指定怪物
	TaskOptionType_UserLevel 				// 角色达到某等级
	TaskOptionType_TargetSlotLevel 			// 指定插槽达到某等级
	TaskOptionType_SlotLevelCount			// 指定数量插槽都达到某等级
	TaskOptionType_CraftSkillLevel			// 指定合成技能达到某等级
	TaskOptionType_RecipeUse				// 使用指定的配方合成
	TaskOptionType_TaskTypeCount			// 完成若干数量的指定类型任务
	TaskOptionType_TargetPosition			// 到达指定坐标点指定半径范围内的区域
*/

func MakeTaskXlsxData() {
	// -------------------- task 1000X --------------------------------------------
	MakeTask1000N()
	// -------------------- task 2000X --------------------------------------------
	MakeTask2000N()
	// -------------------- task 3000X --------------------------------------------
	MakeTask3000N()
}

func MakeTask1000N() {
	idBegin, idEnd := int32(10001), int32(10003)

	rowOptions := make(map[int32][]xlsxTable.TaskXlsxRowOption)
	rowOptions[10001] = []xlsxTable.TaskXlsxRowOption{
		{
			OptionType: proto.TaskOptionType_HandInItem,
			Value:      "1010001,1,100;1010002,1,100;1010003,1,100;1010004,1,100;1010005,1,100;1010006,1,100;3010101,1,5000;3010201,1,5000",
		},
		{
			OptionType: proto.TaskOptionType_KillMonster,
			Value:      "9900002,1,3000;9900003,1,3000;9900004,1,4000",
		},
		{
			OptionType: proto.TaskOptionType_TargetPosition,
			Value:      "111,111,3000;222,222,3000;333,333,4000",
		},
	}
	rowOptions[10002] = rowOptions[10001]
	rowOptions[10003] = rowOptions[10001]

	for id := idBegin; id <= idEnd; id++ {
		row := data.TaskXlsxRow{
			Id:         id,
			Level:      0,
			Name:       "Traveling",
			Decs:       "Go to the %% position",
			Details:    "For the sake of love and peace, please go to the designated place",
			SubSystem:  "1",
			ItemReward: 0,
			ExpReward:  100,        //递增 n + id % 10000 * 增量
			Difficulty: id % 10000, // id % 10000 + 1
			NextTaskId: 0,
			Options:    rowOptions[id],
		}

		// // 指定修改某个id 的某个数据
		// if id == 10002 {
		// 	row.ExpReward = 99
		// 	row.Level = 1
		// }

		AddTaskXlsxData(row)
	}
}

func MakeTask2000N() {
	idBegin, idEnd := int32(20001), int32(20003)

	rowOptions := make(map[int32][]xlsxTable.TaskXlsxRowOption)
	rowOptions[20001] = []xlsxTable.TaskXlsxRowOption{
		{
			OptionType: proto.TaskOptionType_KillMonster,
			Value:      "9900002,1,3000;9900003,1,3000;9900004,1,4000",
		},
	}
	rowOptions[20002] = rowOptions[20001]
	rowOptions[20003] = rowOptions[20001]

	for id := idBegin; id <= idEnd; id++ {
		row := data.TaskXlsxRow{
			Id:         id,
			Level:      0,
			Name:       "Fighting",
			Decs:       "Beat %% Monster %%",
			Details:    "For the sake of love and peace, please defeat those monsters",
			SubSystem:  "1",
			ItemReward: 0,
			ExpReward:  100,        //递增 n + id % 10000 * 增量
			Difficulty: id % 10000, // id % 10000 + 1
			NextTaskId: 0,
			Options:    rowOptions[id],
		}

		// // 指定修改某个id 的某个数据
		// if id == 20002 {
		// 	row.ExpReward = 99
		// 	row.Level = 1
		// }

		AddTaskXlsxData(row)
	}
}

func MakeTask3000N() {
	idBegin, idEnd := int32(30001), int32(30004)

	rowOptions := make(map[int32][]xlsxTable.TaskXlsxRowOption)
	rowOptions[30001] = []xlsxTable.TaskXlsxRowOption{
		{
			OptionType: proto.TaskOptionType_HandInItem,
			Value:      "1010001,1,100;1010002,1,100;1010003,1,100;1010004,1,100;1010005,1,100;1010006,1,100;3010101,1,5000;3010201,1,5000",
		},
	}
	rowOptions[30002] = rowOptions[30001]

	rowOptions[30003] = []xlsxTable.TaskXlsxRowOption{
		{
			OptionType: proto.TaskOptionType_HandInItem,
			Value:      "1010002,1,100;1010003,1,100;1010004,1,100;1010005,1,100;1010006,1,100;3010101,1,5000;3010201,1,5000",
		},
	}
	rowOptions[30004] = rowOptions[30003]

	for id := idBegin; id <= idEnd; id++ {
		row := data.TaskXlsxRow{
			Id:         id,
			Level:      0,
			Name:       "Collecting",
			Decs:       "Collect %% %%",
			Details:    "For love and peace, please give me these",
			SubSystem:  "1,2",
			ItemReward: 0,
			ExpReward:  100,        //递增 n + id % 10000 * 增量
			Difficulty: id % 10000, // id % 10000 + 1
			NextTaskId: 0,
			Options:    rowOptions[id],
		}

		// // 指定修改某个id 的某个数据
		// if id == 30002 {
		// 	row.ExpReward = 99
		// 	row.Level = 1
		// }

		AddTaskXlsxData(row)
	}
}
