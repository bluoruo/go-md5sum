package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

/*
 * 文件 md5 计算工具
 */

// 文件md5值
func fileMd5Sum(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "Open Error:", err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)
	body, err := io.ReadAll(f)
	if err != nil {
		return "io.ReadAll Error:", err
	}
	strMd5sum := fmt.Sprintf("%x", md5.Sum(body))
	return strMd5sum, nil
}

func main() {
	var fileName string
	if len(os.Args) > 1 {
		if strings.Contains(os.Args[1], "-f") {
			flag.StringVar(&fileName, "f", "", "file name")
			flag.Parse()
		} else {
			fmt.Println("by Comanche Lab.	Ver 1.5")
			fileName = os.Args[1]
		}
		strMd5sum, err := fileMd5Sum(fileName)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(filepath.Base(fileName), "- md5sum:", strMd5sum)
	}
}
