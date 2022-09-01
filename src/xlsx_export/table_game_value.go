package xlsx_export

import (
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"gorm.io/gorm"
)

// GameValue.xlsx

var gameValueTableRows = make(map[int32]xlsxTable.GameValueTable)

func ParseGameValue(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}

		setting := xlsxTable.GameValueTable{
			Id:          excel.IntToInt32(row["id"]),
			Value:       excel.IntToInt32(row["value"]),
			StringValue: excel.StringToString(row["strValue"]),
			ValueArray:  excel.IntSliceToString(row["valueArray"]),
			StringArray: excel.StringToString(row["strValueArray"]),
		}

		gameValueTableRows[setting.Id] = setting
	}
	return err
}

func GameValueSaveToDB(db *gorm.DB, curSecUtc time.Time) {
	list := []xlsxTable.GameValueTable{}
	for _, row := range gameValueTableRows {
		row.CreatedAt = curSecUtc
		list = append(list, row)
	}
	WriterToDB(db, curSecUtc, &xlsxTable.GameValueTable{}, len(list), list)
}
