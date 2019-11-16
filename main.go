package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/dfsnotify/model"

	db "github.com/dfsnotify/dal"
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
	option string
)

func init() {
	flag.StringVar(&option, "d", "/data/app/vkgame/public/uploads", "pass directory")
	flag.Parse()
}

func main() {
	// connect mysql
	dsn := "root:1234qwer@tcp(192.168.0.199:3306)/monitor_dir?parseTime=true"
	err := db.InitDB(dsn)
	if err != nil {
		return
	}
	fmt.Println("ok")

	//创建一个监控对象
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watch.Close()
	//添加要监控的对象，文件或文件夹
	err = watch.Add(option)
	if err != nil {
		log.Fatal(err)
	}
	//我们另启一个goroutine来处理监控对象的事件
	go func() {
		nowStr := time.Now().Format("2006/01/02 15:04:05")
		for {
			select {
			case ev := <-watch.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						fmt.Println(nowStr, OP_CREATE, ev.Name)
						record := &model.OpRecord{}
						record.EventTime = nowStr
						record.Operator = OP_CREATE
						record.FilePath = ev.Name
						recordId, err := db.InsertOperatorRecord(record)
						if err != nil {
							err = fmt.Errorf("create failed,insert id: %d,err: %v\n", recordId, err)
							return
						}
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						//fmt.Println(nowStr, OP_WRITE, ev.Name)
						record := &model.OpRecord{}
						record.EventTime = nowStr
						record.Operator = OP_WRITE
						record.FilePath = ev.Name
						recordId, err := db.InsertOperatorRecord(record)
						if err != nil {
							err = fmt.Errorf("write failed,insert id: %d,err: %v\n", recordId, err)
							return
						}
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						//fmt.Println(nowStr, OP_DELETE, ev.Name)
						record := &model.OpRecord{}
						record.EventTime = nowStr
						record.Operator = OP_DELETE
						record.FilePath = ev.Name
						recordId, err := db.InsertOperatorRecord(record)
						if err != nil {
							err = fmt.Errorf("delete failed,insert id: %d,err: %v\n", recordId, err)
							return
						}
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						//fmt.Println(nowStr, OP_RENAME, ev.Name)
						record := &model.OpRecord{}
						record.EventTime = nowStr
						record.Operator = OP_RENAME
						record.FilePath = ev.Name
						recordId, err := db.InsertOperatorRecord(record)
						if err != nil {
							err = fmt.Errorf("rename failed,insert id: %d,err: %v\n", recordId, err)
							return
						}
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						//fmt.Println(nowStr, OP_CHMOD, ev.Name)
						record := &model.OpRecord{}
						record.EventTime = nowStr
						record.Operator = OP_CHMOD
						record.FilePath = ev.Name
						recordId, err := db.InsertOperatorRecord(record)
						if err != nil {
							err = fmt.Errorf("chmod failed,insert id: %d,err: %v\n", recordId, err)
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
