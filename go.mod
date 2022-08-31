module github.com/Meland-Inc/service-xlsx-tool

go 1.17

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gorm.io/driver/mysql v1.3.6 // indirect
	gorm.io/gorm v1.23.8 // indirect
)

// 使用本地go代码仓库方式: https://zhuanlan.zhihu.com/p/109828249
require game-message-core v0.0.0

replace game-message-core => ./game-message-core/messageGo
