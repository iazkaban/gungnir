package entities

type Template struct {
	Name    string    //模板名称
	Modules []*Module //模块列表
}

type Module struct {
	Name        string //名称
	PackagePath string //包的保存路径
}
