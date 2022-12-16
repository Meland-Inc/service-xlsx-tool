package xlsx_export

import (
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"gorm.io/gorm"
)

// SceneArea.xlsx

var sceneAreaRows = make(map[int32]xlsxTable.SceneAreaRow)

func ParseSceneArea(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}

		setting := xlsxTable.SceneAreaRow{
			Id:        excel.IntToInt32(row["id"]),
			Name:      excel.StringToString(row["sceneName"]),
			SceneType: excel.StringToString(row["sceneType"]),
		}
		sceneAreaRows[setting.Id] = setting
	}
	return err
}

func SceneAreaSaveToDB(db *gorm.DB, curSecUtc time.Time) {
	list := []xlsxTable.SceneAreaRow{}
	for _, row := range sceneAreaRows {
		row.CreatedAt = curSecUtc
		list = append(list, row)
	}
	WriterToDB(db, curSecUtc, &xlsxTable.SceneAreaRow{}, len(list), list)
}
