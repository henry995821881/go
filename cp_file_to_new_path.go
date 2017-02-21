package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	ostype = "windows"
)
var listfile []string
var file_type string // xml java html
var file_time string // 2017-01-02
var out_path string  //最后没有斜杠 d:\\ccc
var pwd string

func main() {

	args := os.Args
	file_type = args[1]
	file_time = args[2]
	out_path = args[3]
	pwd, _ = os.Getwd()
	fmt.Println(file_time + "," + file_type + "," + out_path)
	fmt.Println(pwd)

	getFileList(pwd)
	PrintListFileFunc(listfile)
	var result bool
	for _, v := range listfile {

		path_filename := strings.Split(v, ",")

		result, _ = PathExists(out_path + path_filename[0])
		//不存在创建文件夹
		if result != true {
			os.MkdirAll(out_path+path_filename[0], 0777)
		}

		_, err := CopyFile(pwd+path_filename[0]+path_filename[1], out_path+path_filename[0]+path_filename[1])
		if err != nil {
			fmt.Println(err.Error())
		}
	}

}

///
func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

func Filterfunc(path string, f os.FileInfo, err error) error {

	if f == nil {
		return err
	}
	if f.IsDir() {
		return nil
	}

	//转换时间
	tm, _ := time.Parse("2006-01-02", file_time)

	tm1 := f.ModTime()
	//如果时间在设定的之后
	if tm1.Before(tm) {
		return nil
	}

	ok := strings.HasSuffix(f.Name(), "."+file_type)
	if ok {
		fmt.Println(path)
		s1 := strings.Replace(path, pwd, "", -1)
		s2 := strings.Replace(s1, f.Name(), "", -1)

		listfile = append(listfile, s2+","+f.Name())
	}

	return nil

}

func getFileList(path string) string {

	err := filepath.Walk(path, Filterfunc)

	if err != nil {
		fmt.Println("filepath.Walk() returned %v\n", err)
	}

	return " "

}

func PrintListFileFunc(p []string) {

	for index, value := range p {
		fmt.Println("Index = ", index, "Value = ", value)
	}
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
