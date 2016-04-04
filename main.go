// getdir project main.go
/*批量修改文件名称*/
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getFilelist(path string, olds string, news string) { //s是要加的文件名的内容
	var (
		newf string //新文件名
		err  error
	)

	err = filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println("path:", path, "\tfile:", f.Name())
		newf = strings.Replace(path, olds, news+".", -1)
		println("new file:", newf)
		err = os.Rename(path, newf)
		checkErr(err)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	flag.Parse()
	root := flag.Arg(0)
	olds := flag.Arg(1)
	news := flag.Arg(2)
	if flag.NArg() != 3 {
		fmt.Println("修改的目录是：", root)
		fmt.Println("使用的方法是：\n main.exe 要修改的目录名称 要替换的字符 替换后的字符")
		fmt.Println("例如：main.exe d:\\doc . (机要文档).")
		fmt.Println("这样就把\"d:\\doc\\目录.doc\"文件改成了\"d:\\doc\\目录(机要文档).doc\"")
		os.Exit(1)
	}
	getFilelist(root, olds, news)

}
