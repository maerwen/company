package vo

// 员工
type Emp struct {
	EmpNo   int
	EmpName string
	Age     int
	Sex     bool
}

// 部门
type Dept struct {
	DeptNo   int
	DeptName string
	Address  int
}

// 发往页面的文件信息
type File struct {
	Path    string
	Name    string
	Size    int64
	ModTime string
	IsDir   bool
}

// 文件系统页面所需数据
type FileSystem struct {
	Path    string
	FileMap map[int]File
}
