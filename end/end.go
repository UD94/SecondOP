package end

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type file_end interface {
	Read_file()
	Testfunc()
	DeleteFile()
	write_result()
	checkFileIsExist()
	Equal()
}

type F_CONTROL struct {
}

func Testfunc() {
	fmt.Printf("s")
}

func (f_CONTROL F_CONTROL) Read_file(filename string, c chan string) {

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

func (f_CONTROL F_CONTROL) DeleteFile(filename string) {
	//源文件路径

	err := os.Remove(filename) //删除文件test.txt

	if err != nil {

		//如果删除失败则输出 file remove Error!

		fmt.Println("file remove Error!")

		//输出错误详细信息

		fmt.Printf("%s", err)

	} else {

		//如果删除成功则输出 file remove OK!

		fmt.Print("file remove OK!")

	}
}

func (f_CONTROL F_CONTROL) write_result(wireteString string) {

	var filename = "./test.txt"
	var f *os.File
	var err1 error
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Fprintf(os.Stderr, "Err: %s", err1.Error())
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Fprintf(os.Stderr, "Err: %s", err1.Error())
	}
	defer f.Close()
	n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("写入 %d 个字节n", n)
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func (f_CONTROL F_CONTROL) Equal(a, b []string) bool {
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
