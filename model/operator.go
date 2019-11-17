package model

import "time"

// OpRecord 记录目录变化的操作
type OpRecord struct {
	Id        int64     `db:"id"`         //自增id
	EventTime time.Time `db:"event_time"` //事件记录事件
	Operator  string    `db:"operator"`   //事件具体操作
	FilePath  string    `db:"file_path"`  //事件对应的文件
}
