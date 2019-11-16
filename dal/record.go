package db

import (
	"fmt"

	"github.com/learngo/inotify_directory/model"
)

func InsertOperatorRecord(record *model.OpRecord) (recordId int64, err error) {
	if record == nil {
		err = fmt.Errorf("invalid article parameter")
		return
	}
	sqlStr := `insert into record(event_time,operator,file_path) values(?,?,?)`
	result, err := DB.Exec(sqlStr, record.EventTime, record.Operator, record.FilePath)
	if err != nil {
		err = fmt.Errorf("insert record failed, err: %v\n", err)
		return
	}

	recordId, err = result.LastInsertId()
	return
}
