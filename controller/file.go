package controller

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"../vo"
)

var (
	fileFuncMap = make(map[string]interface{})
	fileMap     = make(map[int]vo.File)
	count       int
)

func File(w http.ResponseWriter, r *http.Request) {
	// .是根目录
	// 本地主目录
	root := "."
	saveFileInfos(root)
	// 模板解析
	buff, err := ioutil.ReadFile("src/templates/file/index.html")
	CommonError(w, err)
	// 注册自定义函数It must be called before the template is parsed.
	t, err := template.New("t").Funcs(template.FuncMap{"contains": strings.Contains}).Parse(string(buff))
	TemplateParseError(w, err)
	result := vo.FileSystem{
		Path:    root,
		FileMap: fileMap,
	}
	t.Execute(w, result)
}

// 遍历指定目录或文件夹并执行文件信息存储
func saveFileInfos(root string) {
	err := filepath.Walk(root, walkHandle)
	if err != nil {
		log.Fatalf("遍历根目录出错：\t%s", err.Error())
	}
}

// 根目录遍历与文件信息存储
func walkHandle(path string, info os.FileInfo, err error) error {
	count++
	if info == nil {
		return errors.New("nil")
	}
	fileMap[count] = vo.File{
		Path:    path,
		Name:    info.Name(),
		Size:    info.Size(),
		ModTime: info.ModTime().Format("2006-01-02   3:04:05"),
		IsDir:   info.IsDir(),
	}
	return nil
	/* if strings.Contains(path, "/") {
		count := len(strings.Split(path, "/"))
		for i := 1; i < count; i++ {
			fmt.Print("\t")
		}
	}
	filename := info.Name()
	fmt.Println(filename) */

	// 文件夹操作
	/* if info.IsDir() {
		fmt.Println(filename)
		return nil
	} */
	// 隐藏文件显示后，文件名不包含起始的"."符号，strings.Index(filename, ".")得出-1
	// 如何避免打印隐藏文件呢？
	/* if strings.Index(filename, ".") == -1 {
		return nil
	}
	if len(strings.Split(filename, ".")) > 2 {
		return nil
	} */
}

// 文件名存储
