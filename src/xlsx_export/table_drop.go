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
		// 61000029,3000,1;61000025,5000,1;61000030,1000,1;71010002,120,1

		dropList := xlsxTable.DropList{}

		dropData, ok := row["dropList"].([][]int)
		if !ok {
			err = fmt.Errorf(" Drop.xlsx  dropList  not match to [][]int")
			serviceLog.Error(err.Error())
			continue
		}

		for _, drop := range dropData {
			if len(drop) < 3 {
				err = fmt.Errorf(" Drop.xlsx  id[%v] dropList 配置错误", setting.DropId)
				serviceLog.Error(err.Error())
				continue
			}
			cid := int32(drop[0])
			poss := int32(drop[1])
			quality := int32(drop[2])
			if cid < 1 || poss < 1 || quality < 1 {
				err = fmt.Errorf(" Drop.xlsx  dropId[%v]  dropList data is invalid", setting.DropId)
				serviceLog.Error(err.Error())
				continue
			}

			dropList.List = append(dropList.List, xlsxTable.DropData{
				ObjectCid:   cid,
				Num:         1,
				Possibility: poss,
				Quality:     quality,
			})
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
			if drop.Num < 1 || drop.ObjectCid < 1 {
				err = fmt.Errorf(" Drop.xlsx  dropId[%v]  drop data [%v]is invalid", row.DropId, drop)
				serviceLog.Error(err.Error())
				continue
			}

			if _, exist := itemTableRows[drop.ObjectCid]; !exist {
				err = fmt.Errorf(" Drop.xlsx  dropId[%v]  drop item[%d] not found", row.DropId, drop.ObjectCid)
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
