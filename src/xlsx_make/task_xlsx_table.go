package xlsx_make

import (
	"game-message-core/proto"
	xlsxTable "game-message-core/xlsxTableData"
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

	TaskOptionType_UseRecipe			// 使用指定的配方合成
	{
		OptionType: proto.TaskOptionType_RecipeUse,
		Value:      "20001,3;20002,2", //配方id,次数;配方id,次数;.....
	},

	TaskOptionType_RecipeUseCount		// 累计合成多少次
	{
		OptionType: proto.TaskOptionType_RecipeUseCount,
		Value:      "10", //次数
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
	AddTaskOption(10001,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_KillMonster,
				Value:      "1000001,3",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10002,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_UseItem,
				Value:      "3010101,1",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10003,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_UserLevel,
				Value:      "2",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10004,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_KillMonster,
				Value:      "1000001,10",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10005,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_HandInItem,
				Value:      "4020001,3",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10006,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_UseRecipe,
				Value:      "10301006,1",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10007,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_KillMonster,
				Value:      "1000002,10",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10008,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_SlotLevelCount,
				Value:      "2,4",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10009,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_UseRecipe,
				Value:      "10301002,1",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10010,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_KillMonster,
				Value:      "1000003,10",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10011,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_UserLevel,
				Value:      "5",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10012,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_KillMonster,
				Value:      "1000004,10",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10013,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_SlotLevelCount,
				Value:      "3,4",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10014,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_UserLevel,
				Value:      "8",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10015,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_UseRecipe,
				Value:      "10200001,10",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
	AddTaskOption(10016,
		[]xlsxTable.TaskXlsxOption{
			{
				OptionType: proto.TaskOptionType_KillMonster,
				Value:      "1000005,1",
			},
		},
		[]xlsxTable.TaskXlsxOption{},
	)
}
