package vo

// 员工
type Emp struct {
	empNo   int    `form:"-"`
	empName string `form:"empName"`
	age     int    `form:"age"`
	sex     bool   `form:"sex"`
}

// 部门
type Dept struct {
	deptNo   int    `form:"-"`
	deptName string `form:"deptName"`
	address  int    `form:"address"`
}
