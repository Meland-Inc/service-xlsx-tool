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
			Exp:      excel.IntToInt32(row["exp"]),
		}

		num := 20
		str := ""

		rewardList := xlsxTable.RewardItemList{}
		for i := 1; i <= num; i++ {
			str = fmt.Sprintf("item%dObjectId", i)
			val, ok := row[str]
			if !ok {
				continue
			}

			itemId := excel.IntToInt32(val)
			if 0 == itemId {
				continue
			}

			str = fmt.Sprintf("item%dQuality", i)
			val, ok = row[str]
			if !ok {
				serviceLog.Error("rewardId(%d) 物品(%d)没有配置数量", setting.RewardId, i)
				continue
			}

			itemNum := excel.IntToInt32(val)
			rewardList.List = append(rewardList.List, xlsxTable.RewardItem{
				Cid:      itemId,
				Quantity: itemNum,
			})
		}
		setting.SetRewardList(rewardList)
		rewardTableRows[setting.RewardId] = setting
	}

	return nil
}

func CheckReward() (err error) {
	for _, row := range rewardTableRows {
		rewardList, _ := row.GetRewardList()
		if rewardList == nil || len(rewardList.List) < 1 {
			continue
		}

		for _, reward := range rewardList.List {
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
