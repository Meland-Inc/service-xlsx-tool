package xlsx_export

import (
	"fmt"
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"gorm.io/gorm"
)

// Reward.xlsx

var rewardTableRows = make(map[int32]xlsxTable.RewardTableRow)

func ParseReward(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}
		setting := xlsxTable.RewardTableRow{
			RewardId: excel.IntToInt32(row["id"]),
		}

		// load reward times
		timeData, ok := row["rewardTimes"].([][]int)
		if !ok {
			err = fmt.Errorf(" reward.xlsx  rewardTimes  not match to [][]int")
			serviceLog.Error(err.Error())
			continue
		}
		if len(timeData) < 1 {
			err = fmt.Errorf(" reward.xlsx  rewardTimes   data invalid")
			serviceLog.Error(err.Error())
			continue
		}

		rewardTimeList := xlsxTable.RewardTimeList{}
		for _, v := range timeData {
			if len(v) < 2 {
				err = fmt.Errorf(" reward.xlsx rewardTimes data invalid")
				serviceLog.Error(err.Error())
				continue
			}
			rewardTime := xlsxTable.RewardTime{
				Time:   int32(v[0]),
				Weight: int32(v[1]),
			}
			rewardTimeList.TotalWeight += rewardTime.Weight
			rewardTimeList.RewardTimes = append(rewardTimeList.RewardTimes, rewardTime)
		}
		setting.SetRewardTimeList(rewardTimeList)

		// load reward items
		// 4010001,1,1,100;
		// cid,品质(指定),数量(指定),权重
		rewardItems, ok := row["rewardList"].([][]int)
		if !ok {
			err = fmt.Errorf(" reward.xlsx  rewardList  not match to [][]int")
			serviceLog.Error(err.Error())
			continue
		}

		rewardItemList := xlsxTable.RewardItemList{}
		for _, v := range rewardItems {
			if len(v) < 4 {
				err = fmt.Errorf(" reward.xlsx rewardList data invalid")
				serviceLog.Error(err.Error())
				continue
			}
			rewardItem := xlsxTable.RewardItem{
				Cid:      int32(v[0]),
				Quantity: int32(v[1]),
				Num:      int32(v[2]),
				Weight:   int32(v[3]),
			}
			rewardItemList.TotalWeight += rewardItem.Weight
			rewardItemList.Rewards = append(rewardItemList.Rewards, rewardItem)
		}
		setting.SetRewardList(rewardItemList)

		rewardTableRows[setting.RewardId] = setting
	}

	return err
}

func CheckReward() (err error) {
	for _, row := range rewardTableRows {
		rewardList, _ := row.GetRewardList()
		if rewardList == nil || len(rewardList.Rewards) < 1 {
			continue
		}

		for _, reward := range rewardList.Rewards {
			if 0 == reward.Cid {
				continue
			}
			if _, exist := itemTableRows[reward.Cid]; !exist {
				err = fmt.Errorf("Reward.xlsx id:[%v] reward item [%d] not found", row.RewardId, reward.Cid)
				serviceLog.Error(err.Error())
			}
		}
	}
	return err
}

func RewardSaveToDB(db *gorm.DB, curSecUtc time.Time) {
	list := []xlsxTable.RewardTableRow{}
	for _, Reward := range rewardTableRows {
		Reward.CreatedAt = curSecUtc
		list = append(list, Reward)
	}

	WriterToDB(db, curSecUtc, &xlsxTable.RewardTableRow{}, len(list), list)
}
