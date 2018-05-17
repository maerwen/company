package controller

import (
	"errors"
	"net/http"
	"reflect"
	"strings"
)

// 存储各个方法到map里面,根据uri的一级路径字符串(模块)来调用
var moduleMap = make(map[string]interface{})
var err error

func init() {
	// 模块方法存储
	moduleMap["home"] = Home
	moduleMap["emp"] = Emp
	moduleMap["dept"] = Dept
	moduleMap["office"] = Office
}

func Serve() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe("localhost:8080", nil)
}

// 针对uri为多级路径时调用不同模块的不同方法
// 函数map,函数名称,函数形参列表
/* func Call(funcMap map[string]interface{}, name string, params []interface{}) (resultSet []reflect.Value, err error) {
	f := funcMap[name]
	if f == nil {
		log.Println("err: no such method!")
		return
	}
	fv := reflect.ValueOf(f)
	// // NumIn returns a function type's input parameter count.
	// It panics if the type's Kind is not Func.
	if !(fv.Type().NumIn() == len(params)) {
		err = errors.New("形参列表长度不匹配!")
		return nil, err
	}
	in := make([]reflect.Value, len(params))
	for i, j := range params {
		in[i] = reflect.ValueOf(j)
	}
	return fv.Call(in), nil
} */
func Call(funcMap map[string]interface{}, name string, w http.ResponseWriter, r *http.Request) (resultSet []reflect.Value, err error) {
	f := funcMap[name]
	if f == nil {
		// Error(w ResponseWriter, error string, code int) code是错误状态码
		http.Error(w, "err: no such method!", 401)
		return
	}
	fv := reflect.ValueOf(f)
	// // NumIn returns a function type's input parameter count.
	// It panics if the type's Kind is not Func.
	if !(fv.Type().NumIn() == 2) {
		err = errors.New("形参列表长度不匹配!")
		return nil, err
	}
	in := make([]reflect.Value, 2)
	in[0] = reflect.ValueOf(w)
	in[1] = reflect.ValueOf(r)
	return fv.Call(in), nil
}

// 全局代理方法
func Handler(w http.ResponseWriter, r *http.Request) {
	// 先判断uri中包含几级目录
	// 根目录
	if len(r.RequestURI) == 1 {
		Home(w, r)
		return
	}
	uri := r.RequestURI[1:]
	// 一级目录
	if !strings.Contains(uri, "/") {
		Call(moduleMap, uri, w, r)
	} else {
		// 二级目录
		module := strings.Split(uri, "/")[0]
		Call(moduleMap, module, w, r)
	}
}

// 主页
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to home!"))
	return
}
