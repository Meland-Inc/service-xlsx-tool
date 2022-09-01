package xlsx_export

import (
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"gorm.io/gorm"
)

// Chat.xlsx

var chatTableRows = make(map[xlsxTable.ChatChannelType]xlsxTable.ChatTableRow)

func ParseChat(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}
		setting := xlsxTable.ChatTableRow{
			ChatType: xlsxTable.ChatChannelType(excel.IntToInt32(row["id"])),
			Cd:       excel.IntToInt32(row["talkCD"]),
		}

		chatTableRows[setting.ChatType] = setting
	}

	return nil
}

func CheckChat() (err error) {

	return err
}

func ChatSaveToDB(db *gorm.DB, curSecUtc time.Time) {
	list := []xlsxTable.ChatTableRow{}
	for _, row := range chatTableRows {
		row.CreatedAt = curSecUtc
		list = append(list, row)
	}
	WriterToDB(db, curSecUtc, &xlsxTable.ChatTableRow{}, len(list), list)
}
