package main

import (
	"fmt"
	"os"
)

func main() {

	//当前目录
	// dir, _ := os.Getwd()
	// fmt.Println("pwd：", dir)

	// path := os.Getenv("JAVA_HOME")
	// fmt.Println("path: ", path)

	//data := os.Environ()
	//fmt.Println("system path", data)

	//mkdir
	/*var separator string
	if os.IsPathSeparator('\\') {
		separator = "\\"
	} else {

		separator = "/"
	}

	err := os.Mkdir(dir+separator+"tmp", os.ModePerm)

	if err != nil {
		fmt.Println("mkdir faild", err)
	}*/

	//remove
	/*err := os.MkdirAll("./a", os.ModePerm)
	os.Chdir("./a")
	os.Create("file.txt")
	fmt.Println(err)
	os.Chdir("../")*/

	//err := os.RemoveAll("./a")
	//fmt.Println(err)

	//file1
	// r, w, _ := os.Pipe()
	// w.WriteString("hellow,woow")
	// var s = make([]byte, 20)
	// n, _ := r.Read(s)
	// fmt.Println(string(s[:n]))

	//f, _ := os.Open("tmp.txt")
	//err := f.Chdir()

	//if err != nil {
	//fmt.Println(err) // not is a directory
	//}

	/*f, _ = os.Open("develop")
	err = f.Chdir()

	if err != nil {
		fmt.Println(err)
	}
	*/

	//file2
	/*os.Create("./test.txt")
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
	}

	filelice, err := file.Readdir(5)

	if err != nil {
		fmt.Println(err, filelice)
	}*/

	/*for _, f := range filelice {
		fmt.Println(f.Name())
	}*/

	// file3

	/*s := make([]byte, 10)
	file, _ := os.Open("test.txt")
	defer file.Close()
	file.Seek(-12, 2)
	n, _ := file.Read(s)
	fmt.Println(string(s[:n]))*/

	/////////////////file

}
