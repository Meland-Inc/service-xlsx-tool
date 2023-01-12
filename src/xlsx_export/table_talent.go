package xlsx_export

import (
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"gorm.io/gorm"
)

// TalentTree.xlsx

var talentTableRows = make(map[int32]xlsxTable.TalentTreeRow)

func ParseTalentTree(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}

		setting := xlsxTable.TalentTreeRow{
			NodeId:     excel.IntToInt32(row["id"]),
			TalentType: excel.IntToInt32(row["type"]),
			IsTrunk:    excel.BoolToBool(row["isTrunk"]),
			Layer:      excel.IntToInt32(row["layer"]),
			LvLimit:    excel.IntToInt32(row["lvLimit"]),
		}
		upgradeData := xlsxTable.TalentUpgradeData{
			TalentTreeLv:  excel.IntToInt32(row["upgradeRequireTreeLv"]),
			TalentExpType: setting.TalentType,
			TalentExpList: excel.IntSliceToInt32Slice(row["upgradeEXP"]),
		}
		preNodes := excel.IntSliceToInt32Slice(row["preNode"])
		for _, preNode := range preNodes {
			if preNode <= 0 {
				continue
			}
			upgradeData.PreNodes = append(upgradeData.PreNodes, preNode)
		}

		setting.SetUpgradeData(&upgradeData)
		talentTableRows[setting.NodeId] = setting
	}
	return err
}

func TalentTreeSaveToDB(db *gorm.DB, curSecUtc time.Time) {
	rows := []xlsxTable.TalentTreeRow{}
	for _, row := range talentTableRows {
		row.CreatedAt = curSecUtc
		rows = append(rows, row)
	}
	WriterToDB(db, curSecUtc, &xlsxTable.TalentTreeRow{}, len(rows), rows)
}
