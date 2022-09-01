package xlsx_export

import (
	"fmt"
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"gorm.io/gorm"
)

// buff.xlsx

var buffTableRows = make(map[int32]xlsxTable.BuffTableRow)

func ParseBuff(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}
		setting := xlsxTable.BuffTableRow{
			BuffId:          excel.IntToInt32(row["id"]),
			EffectType:      xlsxTable.BuffEffectType(excel.IntToInt32(row["buffEffect"])),
			GroupId:         excel.IntToInt32(row["buffGroupId"]),
			GroupPriority:   excel.IntToInt32(row["buffPriority"]),
			TotalTime:       excel.IntToInt32(row["totleTime"]),
			TriggerInterval: excel.IntToInt32(row["triggerInterval"]),
		}
		params := excel.IntSliceToInt32Slice(row["buffPara"])
		setting.SetParams(params)
		buffTableRows[setting.BuffId] = setting
	}

	return nil
}

func CheckBuff() (err error) {
	for _, buff := range buffTableRows {
		if buff.TotalTime < 1 {
			err = fmt.Errorf("buff.xlsx id:[%v] totleTime 配置错误", buff.BuffId)
			serviceLog.Error(err.Error())
		}
		if buff.TriggerInterval < 1 {
			err = fmt.Errorf("buff.xlsx id:[%v] triggerInterval 配置错误", buff.BuffId)
			serviceLog.Error(err.Error())
		}
		if len(buff.ParamStr) < 1 {
			err = fmt.Errorf("buff.xlsx id:[%v] buffPara 配置错误", buff.BuffId)
			serviceLog.Error(err.Error())
		}
	}
	return err
}

func SaveBuffToDB(db *gorm.DB, curSecUtc time.Time) {
	list := []xlsxTable.BuffTableRow{}
	for _, buff := range buffTableRows {
		buff.CreatedAt = curSecUtc
		list = append(list, buff)
	}

	WriterToDB(db, curSecUtc, &xlsxTable.BuffTableRow{}, len(list), list)
}
