package main

import (
	"flag"
	"fmt"

	"github.com/dirmonitor/config"
	"github.com/dirmonitor/dal/db"
	"github.com/dirmonitor/logic"
	"github.com/dirmonitor/model"
)

func init() {
	flag.StringVar(&model.Conf, "f", "./monitor.cnf", "pass config")
	flag.StringVar(&model.Directory, "d", "/data/app/vkgame/public/uploads", "pass directory")
	flag.Parse()
}

func main() {
	// connect mysql
	dsn, err := config.ParseConfig(model.Conf)
	if err != nil {
		err = fmt.Errorf("configuration parse failed")
		return
	}
	err = db.InitDB(dsn)
	if err != nil {
		return
	}
	//逻辑处理
	logic.WatchEvent()

}
