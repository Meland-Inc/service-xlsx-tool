package main

import (
	"fmt"

	"github.com/Meland-Inc/service-xlsx-tool/src/common/serviceLog"
)

func main() {
	fmt.Println("this is meland service xlsx export tool.")
	serviceLog.Init(10001, true)

	serviceLog.Error("这里是测试错误日志 ---------- ")
}
