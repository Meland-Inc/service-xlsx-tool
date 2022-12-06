package main

import (
	"os"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"github.com/Meland-Inc/service-xlsx-tool/src/xlsx_export"
	"github.com/Meland-Inc/service-xlsx-tool/src/xlsx_make"
)

//-configdir /Users/44alex/work/meland/meland_svn_res/trunk/settings/config/xlsx

func main() {
	serviceLog.Init(10001, true)

	xlsxModel := os.Getenv("XLSX_MODEL")
	switch xlsxModel {
	case "export":
		xlsx_export.ExportXlsx()
	case "make":
		xlsx_make.MakeXlsxTable()
	default:
		serviceLog.Error("invalid xlsx model [%s]", xlsxModel)
	}
}
