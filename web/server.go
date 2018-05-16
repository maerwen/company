package controller

import (
	"errors"
	"log"
	"net/http"
	"reflect"
	"strings"
)

// 存储各个方法到map里面,根据uri的一级路径字符串(模块)来调用
var funcMap1 = make(map[string]interface{})

// 存储各个方法到map里面,根据uri的二级路径字符串来调用
var funcMap2 = make(map[string]interface{})

func init() {
	// 模块方法存储
	funcMap1["home"] = Home
	funcMap1["emps"] = Emps
	funcMap1["depts"] = Depts
	funcMap1["office"] = Office
}

func Serve() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe("localhost:8080", nil)
}

// 针对uri为多级路径时调用不同模块的不同方法
// 函数map,函数名称,函数形参列表
func Call(funcMap map[string]interface{}, name string, params []interface{}) (resultSet []reflect.Value, err error) {
	f := funcMap[name]
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
		Call(funcMap1, uri, []interface{}{w, r})
		return
	}
	// 二级目录
	arr := strings.Split(uri, "/")
	module := arr[0]
	method := arr[1]
	log.Println(module)
	log.Println(method)
}

// 各模块功能
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to home!"))
	return
}
func Emps(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to emps!"))
	return
}
func Depts(w http.ResponseWriter, r *http.Request) {
}
func Office(w http.ResponseWriter, r *http.Request) {
}
