package entities

import "gungnir/models"

type Module struct {
	ID               int64  `xorm:"id pk autoincr"`     //模块ID
	TemplateID       int64  `xorm:"template_id"`        //模板ID
	Index            int    `xorm:"index"`              //排序
	Name             string `xorm:"name"`               //名称
	BeforeScriptPath string `xorm:"before_script_path"` //解压前脚本路径
	AfterScriptPath  string `xorm:"after_script_path"`  //解压后脚本路径
	PackagePath      string `xorm:"package_path"`       //压缩包的保存路径
}

func (Module) TableName() string {
	return "modules"
}

func (m *Module) ToModule() *models.Module {
	return &models.Module{
		ID:               m.ID,
		Index:            m.Index,
		Name:             m.Name,
		BeforeScriptPath: m.BeforeScriptPath,
		AfterScriptPath:  m.AfterScriptPath,
		PackagePath:      m.PackagePath,
	}
}

func FromModuleModel(m *models.Module, template_id int64) *Module {
	return &Module{
		ID:               m.ID,
		TemplateID:       template_id,
		Index:            m.Index,
		Name:             m.Name,
		BeforeScriptPath: m.BeforeScriptPath,
		AfterScriptPath:  m.AfterScriptPath,
		PackagePath:      m.PackagePath,
	}
}
