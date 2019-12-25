package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/dirmonitor/logic"

	"github.com/gin-gonic/gin"

	"github.com/dirmonitor/config"
	"github.com/dirmonitor/dal/db"
	"github.com/dirmonitor/model"
)

func init() {
	flag.StringVar(&model.Conf, "f", "./monitor.cnf", "pass config")
	flag.StringVar(&model.Directory, "d", "/data/app/vkgame/public/uploads", "pass directory")
	flag.Parse()
}

func GetInfoHandler(ctx *gin.Context) {
	RecordList, err := db.QueryOperatorRecord()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "查询记录信息失败",
		})
	}
	ctx.HTML(http.StatusOK, "views/index.html", gin.H{
		"code": 0,
		"data": RecordList,
	})
}

func main() {
	// connect MySQL
	dsn, err := config.ParseConfig(model.Conf)
	if err != nil {
		err = fmt.Errorf("configuration parse failed")
		return
	}
	err = db.InitDB(dsn)
	if err != nil {
		return
	}

	//引入gin框架
	router := gin.Default()
	router.LoadHTMLGlob("views/*")
	router.GET("/", GetInfoHandler)
	go router.Run(":9999")

	//逻辑处理
	logic.WatchEvent()

}
