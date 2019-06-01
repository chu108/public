package tools

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"github.com/xie1xiao1jun/public/mylog"
)

//检查目录是否存在
func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		mylog.Debug(filename + " not exist")
		exist = false
	}
	return exist
}

//创建目录
func BuildDir(abs_dir string) error {
	return os.MkdirAll(path.Dir(abs_dir), os.ModePerm) //生成多级目录
}

//删除文件或文件夹
func DeleteFile(abs_dir string) error {
	return os.RemoveAll(abs_dir)
}

//获取目录所有文件夹
func GetPathDirs(abs_dir string) (re []string) {
	if CheckFileIsExist(abs_dir) {
		files, _ := ioutil.ReadDir(abs_dir)
		for _, f := range files {
			if f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

//获取目录所有文件夹
func GetPathFiles(abs_dir string) (re []string) {
	if CheckFileIsExist(abs_dir) {
		files, _ := ioutil.ReadDir(abs_dir)
		for _, f := range files {
			if !f.IsDir() {
				re = append(re, f.Name())
			}
		}
	}
	return
}

//获取目录地址
func GetModelPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	// if len(path) > 0 {
	// 	path += "/"
	// }
	path = filepath.Dir(path)
	return path
}

//写入文件
func WriteFile(fname string, src []string, isClear bool) bool {
	BuildDir(fname)
	flag := os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	if !isClear {
		flag = os.O_CREATE | os.O_RDWR | os.O_APPEND
	}
	f, err := os.OpenFile(fname, flag, 0666)
	if err != nil {
		mylog.Error(err)
		return false
	}
	defer f.Close()

	for _, v := range src {
		f.WriteString(v)
		f.WriteString("\r\n")
	}

	return true
}

//读取文件
func ReadFile(fname string) (src []string) {
	f, err := os.OpenFile(fname, os.O_RDONLY, 0666)
	if err != nil {
		return []string{}
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		src = append(src, line)
	}

	return src
}
