package xlsx_export

import (
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"gorm.io/gorm"
)

// SlotLv.xlsx

var SlotLvTableRows = make(map[int32]xlsxTable.SlotLvTableRow)

func ParseSlotLv(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}
		setting := xlsxTable.SlotLvTableRow{
			Position:   excel.IntToInt32(row["slot"]),
			Lv:         excel.IntToInt32(row["lv"]),
			UpExp:      excel.IntToInt32(row["exp"]),
			UpMeld:     excel.IntToInt32(row["useMELD"]),
			HpLimit:    excel.IntToInt32(row["hp"]),
			HpRecovery: excel.IntToInt32(row["hpRecovery"]),
			Att:        excel.IntToInt32(row["att"]),
			AttSpeed:   excel.IntToInt32(row["attSpd"]),
			Def:        excel.IntToInt32(row["def"]),
			CritRate:   excel.IntToInt32(row["critRate"]),
			CritDmg:    excel.IntToInt32(row["critDmg"]),
			HitRate:    excel.IntToInt32(row["hitPoint"]),
			MissRate:   excel.IntToInt32(row["missPoint"]),
			MoveSpeed:  excel.IntToInt32(row["moveSpeed"]),
		}

		SlotLvTableRows[setting.Lv] = setting
	}

	return nil
}

func SlotLvSaveToDB(db *gorm.DB, curSecUtc time.Time) {
	list := []xlsxTable.SlotLvTableRow{}
	for _, SlotLv := range SlotLvTableRows {
		SlotLv.CreatedAt = curSecUtc
		list = append(list, SlotLv)
	}

	WriterToDB(db, curSecUtc, &xlsxTable.SlotLvTableRow{}, len(list), list)
}
