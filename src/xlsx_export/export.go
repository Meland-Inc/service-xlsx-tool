package xlsx_export

import (
	"os"
	"time"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
)

func ExportXlsx() {
	configDir := os.Getenv("XLSX_DIR")
	if configDir == "" {
		panic("config dir is empty")
	}

	InitTables()

	err := ParseTables(configDir)
	if err != nil {
		panic(err)
	}

	err = SaveToDB()
	if err != nil {
		panic(err)
	}
}

func InitTables() {
	// RegisterTable("",)
}

func ParseTables(configDir string) (err error) {
	for _, parse := range GetParseTables() {
		configFilePath, checkErr := checkConfigDir(configDir, parse.TableName)
		if checkErr != nil {
			err = checkErr
			serviceLog.Error(err.Error())
			continue
		}

		rows, tableErr := loadXlsxData(configFilePath, parse.SheetName)
		if tableErr != nil {
			err = tableErr
			serviceLog.Error(err.Error())
			continue
		}

		if parseErr := parse.ParseSheetF(rows); parseErr != nil {
			err = parseErr
			serviceLog.Error(err.Error())
		}
	}
	return err
}

func CheckTablesError() (err error) {
	for _, parse := range GetParseTables() {
		if parse.CheckF == nil {
			continue
		}

		if checkErr := parse.CheckF(); checkErr != nil {
			err = checkErr
			serviceLog.Error(err.Error())
		}
	}
	return err
}

func SaveToDB() (err error) {
	db, err := GetTableDB(nil) // TODO ... XLSX DATA MODULE
	if err != nil {
		serviceLog.Error(err.Error())
		return err
	}

	curUtc := time.Now().UTC()
	for _, parse := range GetParseTables() {
		if parse.WriteF == nil {
			continue
		}

		if writeErr := parse.WriteF(db, curUtc); writeErr != nil {
			err = writeErr
			serviceLog.Error(err.Error())
		}
	}
	return err
}
