package xlsx_export

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Meland-Inc/service-xlsx-tool/src/common/excel"
	"github.com/Meland-Inc/service-xlsx-tool/src/common/gormDB"
	"gorm.io/gorm"
)

type ParseSheetFunc func(rows []map[string]interface{}) (err error)

type XlsxParse struct {
	TableName   string
	SheetName   string
	ParseSheetF ParseSheetFunc
	CheckF      func() error
	WriteF      func(db *gorm.DB, curSecUtc time.Time) error
}

var parseTables = make([]XlsxParse, 0, 0)

func GetParseTables() []XlsxParse {
	return parseTables
}

func RegisterTable(
	tableName string,
	parseF ParseSheetFunc,
	checkF func() error,
	writeF func(db *gorm.DB, curSecUtc time.Time) error,
) {
	if tableName == "" || parseF == nil {
		panic("invalid tableName || parseF  is nill")
	}

	parseTables = append(parseTables, XlsxParse{
		TableName:   tableName,
		ParseSheetF: parseF,
		CheckF:      checkF,
		WriteF:      writeF,
	})
}

func GetTableDB(models []interface{}) (*gorm.DB, error) {
	host := os.Getenv("MELAND_CONFIG_DB_HOST")
	port := os.Getenv("MELAND_CONFIG_DB_PORT")
	user := os.Getenv("MELAND_CONFIG_DB_USER")
	pass := os.Getenv("MELAND_CONFIG_DB_PASS")
	dbName := os.Getenv("MELAND_CONFIG_DB_DATABASE")
	return gormDB.InitGormDB(host, port, user, pass, dbName, models)
}

func WriterToDB(db *gorm.DB, curUtc time.Time, pTypes interface{}, dataLength int, datas interface{}) {
	err := db.Transaction(func(tx *gorm.DB) error {
		// 清空老的配置数据
		err := tx.Where("created_at < ?", curUtc).Delete(pTypes).Error
		if err != nil {
			return err
		}

		err = tx.Model(pTypes).CreateInBatches(datas, dataLength).Error
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func checkConfigDir(configDir, filename string) (configFilePath string, err error) {
	dirInfo, err := os.Stat(configDir)
	if err != nil {
		return
	}

	if !dirInfo.IsDir() {
		err = fmt.Errorf("%s is not a directory", configDir)
		return
	}

	configFilePath = filepath.Join(configDir, filename)

	configInfo, err := os.Stat(configFilePath)
	if err != nil {
		return
	}

	if !configInfo.Mode().IsRegular() {
		err = fmt.Errorf("config %s not found", configFilePath)
	}

	return
}

func loadXlsxData(configFilePath, sheetName string) ([]map[string]interface{}, error) {
	ef, err := excelize.OpenFile(configFilePath)
	if err != nil {
		return nil, err
	}

	if 0 == len(sheetName) {
		sheetName = excel.FirstSheet(ef)
	}

	return excel.ParseExcelSheet(ef, sheetName)
}
