package model

import "time"

// OpRecord 记录目录变化的操作
type OpRecord struct {
	Id        int64     `db:"id"`         //自增id
	EventTime time.Time `db:"event_time"` //事件时间
	Operator  string    `db:"operator"`   //具体操作
	FilePath  string    `db:"file_path"`  //对应文件
}
