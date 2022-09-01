package xlsx_export

import (
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"gorm.io/gorm"
)

// RoleLv.xlsx

var roleLvTableRows = make(map[int32]xlsxTable.RoleLvTableRow)

func ParseRoleLv(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}
		setting := xlsxTable.RoleLvTableRow{
			Lv:           excel.IntToInt32(row["lv"]),
			Exp:          excel.IntToInt32(row["exp"]),
			HpLimit:      excel.IntToInt32(row["hp"]),
			HpRecovery:   excel.IntToInt32(row["hpRecovery"]),
			DeathExpLoss: excel.IntToInt32(row["deathExpLoss"]),
			Att:          excel.IntToInt32(row["att"]),
			AttSpeed:     excel.IntToInt32(row["attSpd"]),
			Def:          excel.IntToInt32(row["def"]),
			CritRate:     excel.IntToInt32(row["critRate"]),
			CritDmg:      excel.IntToInt32(row["critDmg"]),
			HitRate:      excel.IntToInt32(row["hitPoint"]),
			MissRate:     excel.IntToInt32(row["missPoint"]),
			MoveSpeed:    excel.IntToInt32(row["moveSpeed"]),
		}

		roleLvTableRows[setting.Lv] = setting
	}

	return nil
}

func RoleLvSaveToDB(db *gorm.DB, curSecUtc time.Time) {
	list := []xlsxTable.RoleLvTableRow{}
	for _, RoleLv := range roleLvTableRows {
		RoleLv.CreatedAt = curSecUtc
		list = append(list, RoleLv)
	}

	WriterToDB(db, curSecUtc, &xlsxTable.RoleLvTableRow{}, len(list), list)
}
