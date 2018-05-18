package controller

// 错误处理
import (
	"fmt"
	"log"
	"net/http"
)

// 模板解析出错处理,出错返回false
func TemplateParseError(w http.ResponseWriter, err error) bool {
	if err != nil {
		log.Printf("parse template error: %s!\n", err.Error())
		fmt.Fprintf(w, "parse template error: %s!", err.Error())
	}
	return err == nil
}

// 一般错误处理,出错返回false
func CommonError(w http.ResponseWriter, err error) bool {
	if err != nil {
		log.Printf("error: %s!\n", err.Error())
		fmt.Fprintf(w, "error: %s!", err.Error())
	}
	return err == nil
}
