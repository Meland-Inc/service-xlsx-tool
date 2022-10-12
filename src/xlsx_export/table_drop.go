package xlsx_export

import (
	"fmt"
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"gorm.io/gorm"
)

// Drop.xlsx

var dropTableRows = make(map[int32]xlsxTable.DropTableRow)

func ParseDrop(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}

		setting := xlsxTable.DropTableRow{
			DropId: excel.IntToInt32(row["id"]),
		}

		// 1010001,	   1,		1,		20;
		// cid,		品质(指定),数量(指定) 万份比例,

		dropList := xlsxTable.DropList{}
		dropData, ok := row["dropList"].([][]int)
		if !ok {
			err = fmt.Errorf(" Drop.xlsx  dropList  not match to [][]int")
			serviceLog.Error(err.Error())
			continue
		}

		for _, drop := range dropData {
			if len(drop) < 4 {
				err = fmt.Errorf(" Drop.xlsx  id[%v] dropList 配置错误", setting.DropId)
				serviceLog.Error(err.Error())
				continue
			}
			dropData := xlsxTable.DropData{
				Cid:         int32(drop[0]),
				Quality:     int32(drop[1]),
				Num:         int32(drop[2]),
				Possibility: int32(drop[3]),
			}

			dropList.List = append(dropList.List, dropData)
		}

		setting.SetDropList(dropList)
		dropTableRows[setting.DropId] = setting
	}

	return err
}

func CheckDrop() (err error) {
	for _, row := range dropTableRows {
		dropList, _ := row.GetDropList()
		if dropList == nil || len(dropList.List) == 0 {
			err = fmt.Errorf(" Drop.xlsx  dropId[%v]  dropList data is invalid", row.DropId)
			serviceLog.Error(err.Error())
			continue
		}

		for _, drop := range dropList.List {
			if drop.Num < 1 || drop.Cid < 1 || drop.Possibility < 1 || drop.Quality < 1 {
				err = fmt.Errorf(" Drop.xlsx  dropId[%v]  drop data [%v]is invalid", row.DropId, drop)
				serviceLog.Error(err.Error())
				continue
			}

			if _, exist := itemTableRows[drop.Cid]; !exist {
				err = fmt.Errorf(" Drop.xlsx  dropId[%v]  drop item[%d] not found", row.DropId, drop.Cid)
				serviceLog.Error(err.Error())
			}
		}
	}

	return err
}

func DropSaveToDB(db *gorm.DB, curSecUtc time.Time) {
	list := []xlsxTable.DropTableRow{}
	for _, row := range dropTableRows {
		row.CreatedAt = curSecUtc
		list = append(list, row)
	}
	WriterToDB(db, curSecUtc, &xlsxTable.DropTableRow{}, len(list), list)
}
