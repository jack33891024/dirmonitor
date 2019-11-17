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
	sqlStr := `insert into record(operator,file_path) values(?,?)`
	result, err := DB.Exec(sqlStr, record.Operator, record.FilePath)
	if err != nil {
		err = fmt.Errorf("Execute SQL failed, err: %v\n", err)
		return
	}
	_, err = result.LastInsertId()
	if err != nil {
		err = fmt.Errorf("insert record failed, err: %v\n", err)
		return
	}
	return
}
