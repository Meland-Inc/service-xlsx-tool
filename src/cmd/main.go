package main

import (
	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
	"github.com/Meland-Inc/service-xlsx-tool/src/xlsx_export"
)

//-configdir /Users/44alex/work/meland/meland_svn_res/trunk/settings/config/xlsx

func main() {
	serviceLog.Init(10001, true)
	xlsx_export.ExportXlsx()
}
