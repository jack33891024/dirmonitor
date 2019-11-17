package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dirmonitor/config"
	"github.com/dirmonitor/dal/db"
	"github.com/dirmonitor/model"
	"github.com/fsnotify/fsnotify"
)

const (
	OP_CREATE = "创建文件"
	OP_WRITE  = "写入文件"
	OP_DELETE = "删除文件"
	OP_RENAME = "重命名文件"
	OP_CHMOD  = "修改权限"
)

var (
	conf      string //cmd传的配置文件
	directory string //cmd传的目录
)

func init() {
	flag.StringVar(&conf, "f", "/etc/monitor.cnf", "pass config")
	flag.StringVar(&directory, "d", "/data/app/vkgame/public/uploads", "pass directory")
	flag.Parse()
}

func main() {
	// connect mysql
	dsn, err := config.ParseConfig(conf)
	if err != nil {
		err = fmt.Errorf("configuration parse failed")
		return
	}
	err = db.InitDB(dsn)
	if err != nil {
		return
	}

	//创建一个监控对象
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watch.Close()
	//添加要监控的对象，文件或文件夹
	err = watch.Add(directory)
	if err != nil {
		log.Fatal(err)
	}
	//我们另启一个goroutine来处理监控对象的事件
	go func() {
		for {
			select {
			case ev := <-watch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						record := &model.OpRecord{}
						record.Operator = OP_CREATE
						record.FilePath = ev.Name
						err := db.InsertOperatorRecord(record)
						if err != nil {
							err = fmt.Errorf("create failed,err: %v\n", err)
							return
						}
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						//fmt.Println(nowStr, OP_WRITE, ev.Name)
						record := &model.OpRecord{}
						record.Operator = OP_WRITE
						record.FilePath = ev.Name
						err := db.InsertOperatorRecord(record)
						if err != nil {
							err = fmt.Errorf("write failed,err: %v\n", err)
							return
						}
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						//fmt.Println(nowStr, OP_DELETE, ev.Name)
						record := &model.OpRecord{}
						record.Operator = OP_DELETE
						record.FilePath = ev.Name
						err := db.InsertOperatorRecord(record)
						if err != nil {
							err = fmt.Errorf("delete failed,err: %v\n", err)
							return
						}
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						//fmt.Println(nowStr, OP_RENAME, ev.Name)
						record := &model.OpRecord{}
						record.Operator = OP_RENAME
						record.FilePath = ev.Name
						err := db.InsertOperatorRecord(record)
						if err != nil {
							err = fmt.Errorf("rename failederr: %v\n", err)
							return
						}
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						//fmt.Println(nowStr, OP_CHMOD, ev.Name)
						record := &model.OpRecord{}
						record.Operator = OP_CHMOD
						record.FilePath = ev.Name
						err := db.InsertOperatorRecord(record)
						if err != nil {
							err = fmt.Errorf("chmod failed,err: %v\n", err)
							return
						}
					}
				}
			case err := <-watch.Errors:
				{
					log.Println("error : ", err)
					return
				}
			}
		}
	}()
	//循环
	select {}
}
