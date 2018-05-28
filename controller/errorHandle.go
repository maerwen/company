package controller

// 错误处理
import (
	"fmt"
	"log"
	"net/http"
)

// 模板解析出错处理,出错返回false
func TemplateParseError(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Fprintf(w, "parse template error: %s!", err.Error())
		log.Fatalf("parse template error: %s!\n", err.Error())
	}
}

// 一般错误处理,出错返回false
func CommonError(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Fprintf(w, "error: %s!", err.Error())
		log.Fatalf("error: %s!\n", err.Error())
	}
}
