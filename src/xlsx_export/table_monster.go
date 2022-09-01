package xlsx_export

import (
	"fmt"
	"time"

	xlsxTable "game-message-core/xlsxTableData"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"gorm.io/gorm"
)

// Monster.xlsx

var monsterTableRows = make(map[int32]xlsxTable.MonsterTableRow)

func ParseMonster(rows []map[string]interface{}) (err error) {
	for _, row := range rows {
		if row["id"] == "" || row["id"] == "0" {
			continue
		}
		setting := xlsxTable.MonsterTableRow{
			Cid:             excel.IntToInt32(row["id"]),
			Name:            excel.StringToString(row["name"]),
			BodyRadius:      excel.IntToInt32(row["bodyCapacity"]),
			AttType:         xlsxTable.MonsterAttType(excel.IntToInt32(row["attType"])),
			LockEnemyRadius: excel.IntToInt32(row["lockEnemyRange"]),
			CombatDist:      excel.IntToInt32(row["combatDist"]),
			DropId:          excel.IntToInt32(row["dropId"]),
			SkillSequence:   excel.IntSliceToString(row["skillSequence"]),
			Att:             excel.IntToInt32(row["att"]),
			AttSpeed:        excel.IntToInt32(row["attSpd"]),
			Def:             excel.IntToInt32(row["def"]),
			HpLimit:         excel.IntToInt32(row["hp"]),
			CritRate:        excel.IntToInt32(row["critRate"]),
			CritDmg:         excel.IntToInt32(row["critDmg"]),
			HitRate:         excel.IntToInt32(row["hitPoint"]),
			MissRate:        excel.IntToInt32(row["missPoint"]),
			MoveSpeed:       excel.IntToInt32(row["moveSpeed"]),
			PushDmg:         excel.IntToInt32(row["pushDmg"]),
			PushDist:        excel.IntToInt32(row["pushDist"]),
		}

		monsterTableRows[setting.Cid] = setting
	}

	return nil
}

func CheckMonster() (err error) {
	for _, row := range monsterTableRows {
		if _, exist := dropTableRows[row.DropId]; !exist {
			err = fmt.Errorf("Monster.xlsx id:[%v] dropId[%d] not found", row.Cid, row.DropId)
			serviceLog.Error(err.Error())
		}
	}
	return err
}

func MonsterSaveToDB(db *gorm.DB, curSecUtc time.Time) {
	list := []xlsxTable.MonsterTableRow{}
	for _, Monster := range monsterTableRows {
		Monster.CreatedAt = curSecUtc
		list = append(list, Monster)
	}

	WriterToDB(db, curSecUtc, &xlsxTable.MonsterTableRow{}, len(list), list)
}
