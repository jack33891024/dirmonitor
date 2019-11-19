package db

import (
	"fmt"

	"github.com/dirmonitor/model"
)

// InsertOperatorRecord 写入记录信息到数据库
func InsertOperatorRecord(record *model.OpRecord) (err error) {
	if record == nil {
		err = fmt.Errorf("invalid article parameter")
		return
	}
	sqlStr := `insert into 
					record(operator,file_path) 
				values(?,?)`
	result, err := DB.Exec(sqlStr, record.Operator, record.FilePath)
	if err != nil {
		err = fmt.Errorf("Execute SQL failed,err: %v\n", err)
		return
	}
	_, err = result.LastInsertId()
	if err != nil {
		err = fmt.Errorf("insert record failed, err: %v\n", err)
		return
	}
	return
}

// 查询展示最近50条
func QueryOperatorRecord() (RecordList []*model.OpRecord, err error) {
	sqlStr := `select 
					id,event_time,operator,file_path 
       			from 
					record
				order by 
					event_time desc
				limit 
					25`

	err = DB.Select(&RecordList, sqlStr)
	if err != nil {
		err = fmt.Errorf("查询记录信息失败\n")
		return
	}

	return
}

// 根据时间区间查询
