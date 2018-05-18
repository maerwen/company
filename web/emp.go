package controller

import (
	"errors"
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
	empFuncMap["index"] = FindEmps
	empFuncMap["insert"] = InsertEmp
	empFuncMap["update"] = UpdateEmp
	empFuncMap["delete"] = DeleteEmp
}

// 对模块下面各个方法进行代理
func Emp(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.RequestURI[1:], "/") {
		FindEmps(w, r)
		return
	}
	path := r.RequestURI[1:]
	if strings.Contains(path, "?") {
		path = strings.Split(r.RequestURI[1:], "?")[0]
	}
	method := strings.Split(path, "/")[1]
	_, err = Call(empFuncMap, method, w, r)
	if !CommonError(w, err) {
		return
	}
}

// GET	查询所有emp
// POST	根据指定条件查询emp
func FindEmps(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("src/templates/emp/index.html")
	if !TemplateParseError(w, err) {
		return
	}
	err = t.Execute(w, empMap)
	if !CommonError(w, err) {
		return
	}
}

// GET	显示新增emp页面
// POST	插入一条emp数据
func InsertEmp(w http.ResponseWriter, r *http.Request) {
	switch r.Method { // r.Method是GET, POST, PUT, etc..大写!!!
	case "GET": //访问网页
		r.ParseForm()
		t, err := template.ParseFiles("src/templates/emp/insert.html")
		if !TemplateParseError(w, err) {
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
		if !CommonError(w, err) {
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
		log.Println("插入了一条emp数据")
		fallthrough
	default:
		FindEmps(w, r)
	}
}

// GET	显示emp修改页面
// POST	修改一条emp数据
func UpdateEmp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		r.ParseForm()
		empIdArr := r.Form["empNo"]
		if empIdArr == nil {
			err = errors.New("no such parameter")
			if !CommonError(w, err) {
				return
			}
		}
		empId, err := strconv.Atoi(empIdArr[0])
		if !CommonError(w, err) {
			return
		}
		emp := empMap[empId]
		t, err := template.ParseFiles("src/templates/emp/update.html")
		if !TemplateParseError(w, err) {
			return
		}
		err = t.Execute(w, emp)
		if !CommonError(w, err) {
			return
		}
	case "POST":
		r.ParseForm()
		empId, err := strconv.Atoi(r.Form["empNo"][0])
		if !CommonError(w, err) {
			return
		}
		empName := r.Form["empName"][0]
		age, err := strconv.Atoi(r.Form["age"][0])
		if !CommonError(w, err) {
			return
		}
		sexStr := r.Form["sex"][0]
		var sex bool
		if sexStr == "0" {
			sex = false
		} else {
			sex = true
		}
		emp := vo.Emp{
			EmpNo:   empId,
			EmpName: empName,
			Age:     age,
			Sex:     sex,
		}
		empMap[empId] = emp
		log.Println("修改了一条emp数据")
		fallthrough
	default:
		FindEmps(w, r)
	}
}

// GET	删除一条emp数据
func DeleteEmp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		r.ParseForm()
		empIdStr := r.Form["empNo"][0]
		empId, err := strconv.Atoi(empIdStr)
		if !CommonError(w, err) {
			return
		}
		// 从map里面删除
		delete(empMap, empId)
		log.Println("删除了一条emp数据")
		fallthrough
	default:
		FindEmps(w, r)
	}
}
