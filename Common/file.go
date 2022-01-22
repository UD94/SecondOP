package common

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func Equal(a, b []string) bool {
	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func Write_result(wireteString string, filename string) {

	var f *os.File
	var err1 error
	if CheckFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件

	} else {
		f, err1 = os.Create(filename) //创建文件

	}
	defer f.Close()
	n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("写入 %d 个字节n", n)
}

func Read_file(filename string, c chan string) {

	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		c <- line
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
	close(c)
}
