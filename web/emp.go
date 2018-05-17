package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../vo"
)

var empMap = make(map[int]vo.Emp)
var empNo int
var empFuncMap = make(map[string]interface{})

// emp模块方法存储
func init() {
	empFuncMap["insert"] = InsertEmp
	empFuncMap["index"] = FindEmps
}

// 对模块下面各个方法进行代理
func Emp(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.RequestURI[1:], "/") {
		FindEmps(w, r)
		return
	}
	method := strings.Split(r.RequestURI[1:], "/")[1]
	_, err = Call(empFuncMap, method, w, r)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

// 查询并列出所有的emp
func FindEmps(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("src/templates/emp/index.html")
	if !ParseTemplateError(w, err) {
		return
	}
	err = t.Execute(w, empMap)
	if err != nil {
		log.Printf("parse template error: %s!\n", err.Error())
		return
	}
}

// 插入一条数据
func InsertEmp(w http.ResponseWriter, r *http.Request) {
	switch r.Method { // r.Method是GET, POST, PUT, etc..大写!!!
	case "GET": //访问网页
		r.ParseForm()
		t, err := template.ParseFiles("src/templates/emp/create.html")
		if !ParseTemplateError(w, err) {
			return
		}
		t.Execute(w, nil)
	case "POST":
		// 请求实体数据读取
		/* buff := make([]byte, 8192)
		defer r.Body.Close()
		_, err = r.Body.Read(buff)
		fmt.Printf("%s\n", buff) */

		// 表单数据读取,但得到的是数组
		// 		ParseForm已经读取了Request Body里的数据
		// 		func (r *Request) ParseMultipartForm(maxMemory int64) error
		// 		func (r *Request) FormValue(key string) string
		// 		func (r *Request) FormFile(key string)
		// 		它们可能会直接或间接的调用ParseForm，同样会造成Body数据被读取。
		r.ParseForm()
		empNameArr := r.Form["empName"]
		ageArr := r.Form["age"]
		sexArr := r.Form["sex"]
		age, err := strconv.Atoi(ageArr[0])
		if err != nil {
			http.Error(w, "年龄格式不正确", 233)
			return
		}
		empNo++
		// 模拟存储到数据库
		emp := vo.Emp{
			EmpNo:   empNo,
			EmpName: empNameArr[0],
			Age:     age,
			Sex:     sexArr[0] == "1",
		}
		empMap[empNo] = emp
		log.Println("向数据库插入了一条emp数据")
		fallthrough
	default:
		FindEmps(w, r)
	}
}
