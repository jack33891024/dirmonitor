package model

// OpRecord 记录目录变化的操作
type OpRecord struct {
	Id        int64  `db:"id"`
	EventTime string `db:"event_time"`
	Operator  string `db:"operator"`
	FilePath  string `db:"file_path"`
}
