package utils

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func NewDateFilePath(oldFilePath string, dateFileFormat string) string {
	var format string
	if dateFileFormat == "" {
		format = "20060102"
	} else {
		format = dateFileFormat
	}

	paths, fileName := filepath.Split(oldFilePath)
	fileNames := strings.Split(fileName, ".")
	newFileName := fileNames[0] + time.Now().Format(format) + "." + fileNames[1]
	return paths + newFileName
}

func WriteStrToFile(filePath string, content string, module string, log func(log string)) {
	file, err := os.OpenFile(filePath, os.O_CREATE, 0644)
	if err != nil {
		log("创建输入文件对象失败：" + err.Error())
	} else {
		n, _ := file.Seek(0, io.SeekEnd)
		_, err = file.WriteAt([]byte(content+"\n"), n)
		log("输出数据：" + content + "到" + filePath)
	}
	defer file.Close()
}
