package models

type Template struct {
	ID       int64    `json:"id"`       //模板ID
	Name     string   `json:"name"`     //模板名
	Language string   `json:"language"` //适用语言
	Modules  []Module `json:"modules"`  //模块列表
}

type Module struct {
	ID               int64  `json:"id"`                 //模块ID
	Index            int    `json:"index"`              //模块顺序
	Name             string `json:"name"`               //名称
	BeforeScriptPath string `json:"before_script_path"` //解压前脚本路径
	AfterScriptPath  string `json:"after_script_path"`  //解压后脚本路径
	PackagePath      string `json:"package_path"`       //压缩包的保存路径
}
