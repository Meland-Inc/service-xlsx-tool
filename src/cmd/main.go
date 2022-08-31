package main

import (
	"fmt"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
)

//-configdir /Users/44alex/work/meland/meland_svn_res/trunk/settings/config/xlsx

func main() {
	fmt.Println("--------------begin--------------")
	serviceLog.Init(10001, true)
}
