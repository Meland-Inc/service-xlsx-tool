package xlsx_export

import (
	xlsxTable "game-message-core/xlsxTableData"
	"time"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"gorm.io/gorm"
)

// Item.xlsx
var itemTableRows = make(map[int32]xlsxTable.ItemTable)

func ParseItem(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}

		setting := xlsxTable.ItemTable{
			ItemCid:  excel.IntToInt32(row["id"]),
			Name:     excel.StringToString(row["name"]),
			ItemType: excel.IntToInt32(row["type"]),
			UserType: excel.IntToInt32(row["userType"]),
			CanMint:  excel.IntToInt32(row["canMint"]) == 1,
		}

		itemTableRows[setting.ItemCid] = setting
	}

	return err
}

func ItemSaveToDB(db *gorm.DB, curSecUtc time.Time) {
	list := []xlsxTable.ItemTable{}
	for _, row := range itemTableRows {
		row.CreatedAt = curSecUtc
		list = append(list, row)
	}
	WriterToDB(db, curSecUtc, &xlsxTable.ItemTable{}, len(list), list)
}
