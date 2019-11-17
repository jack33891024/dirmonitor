package config

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ParseConfig(fileName string) (dsn string, err error) {
	var user, pass, host, port, database string
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}

	lineArr := strings.Split(string(data), "\r\n")
	for _, line := range lineArr {
		// 去除首尾空行
		line = strings.TrimSpace(line)
		// 忽略空行
		if len(line) == 0 {
			continue
		}
		// 忽略注释
		if line[0] == ';' || line[0] == '#' {
			continue
		}

		index := strings.Index(line, "=")
		if index == -1 {
			err = fmt.Errorf("invalid parameter")
			return
		}
		key := strings.TrimSpace(line[:index])
		value := strings.TrimSpace(line[index+1:])

		// 配置文件取值
		switch {
		case key == "user":
			user = value
		case key == "pass":
			pass = value
		case key == "host":
			host = value
		case key == "port":
			port = value
		case key == "database":
			database = value
		}
	}
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, database)
	return
}
