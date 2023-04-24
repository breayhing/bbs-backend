package main

import (
	"zhongzhu-bbs/model"
	"zhongzhu-bbs/routers"
)

func main() {
	//引用数据库
	model.InitDb()

	routers.InitRouter()

}
