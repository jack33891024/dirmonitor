package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// Open File
	f, err := os.Open("index.php")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Get the content
	contentType, err := GetFileContentType(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content Type: ", contentType)
}

func GetFileContentType(out *os.File) (string, error) {

	// 仅使用前512个字节来检测内容类型。
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	// 使用 net/http 包中的的DectectContentType函数,它将始终返回有效的 MIME 类型
	// 对于没有匹配的未知类型，将返回 "application/octet-stream"
	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
