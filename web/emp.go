package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var empFuncMap = make(map[string]interface{})

func init() {
	empFuncMap["create"] = CreateEmp
	empFuncMap["empHome"] = EmpHome
}

// 模块下面方法代理
func Emp(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.RequestURI[1:], "/") {
		EmpHome(w, r)
		return
	}
	method := strings.Split(r.RequestURI[1:], "/")[1]
	_, err = Call(empFuncMap, method, w, r)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}
func EmpHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to empHome!"))
	return
}
func CreateEmp(w http.ResponseWriter, r *http.Request) bool {
	switch r.Method { // r.Method是GET, POST, PUT, etc..大写!!!
	case "GET": //访问网页
		r.ParseForm()
		t, err := template.ParseFiles("templates/emps/create.html")
		if err != nil {
			fmt.Fprintf(w, "parse template error: %s!", err.Error())
			return false
		}
		t.Execute(w, nil)
	case "POST":
		// 请求实体数据读取
		buff := make([]byte, 8192)
		_, err = r.Body.Read(buff)
		fmt.Printf("%s\n", buff)

		/* 表单数据读取,但得到的是数组 */
		/* r.ParseForm()
			empName := r.Form["empName"]
			age := r.Form["age"]
			sex := r.Form["sex"]
			fmt.Printf("empName:\t%s\n", empName)
			fmt.Printf("age:\t%s\n", age)
			fmt.Printf("sex:\t%s\n", sex)
		default: */
	}
	return true
}
