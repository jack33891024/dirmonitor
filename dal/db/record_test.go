package db

import (
	"testing"

	"github.com/dirmonitor/model"
)

func init() {
	dsn := "root:1234qwer@tcp(192.168.0.199:3306)/monitor_dir?parseTime=true"
	err := InitDB(dsn)
	if err != nil {
		return
	}
}

func TestInsertOperatorRecord(t *testing.T) {
	record := &model.OpRecord{}

	record.Operator = "创建文件"
	record.FilePath = "/data/app/vkgame/public/uploads/status.php"

	err := InsertOperatorRecord(record)
	if err != nil {
		t.Fatalf("insert record failed, err: %v\n", err)
	}

	t.Logf("insert success")
}
