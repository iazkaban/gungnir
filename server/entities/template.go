package entities

import "gungnir/models"

type Template struct {
	ID       int64    `xorm:"id pk autoincr"` //模板ID
	Name     string   `xorm:"name"`           //模板名
	Language string   `xorm:"language"`       //适用语言
	Modules  []Module `xorm:"-"`              //模块列表
}

func (Template) TableName() string {
	return "templates"
}

func (t *Template) ToModel() *models.Template {
	m := &models.Template{
		ID:       t.ID,
		Name:     t.Name,
		Language: t.Language,
	}

	mList := make([]models.Module, 0, len(t.Modules))

	for _, module := range t.Modules {
		newM := models.Module{
			Index:            module.Index,
			Name:             module.Name,
			BeforeScriptPath: module.BeforeScriptPath,
			AfterScriptPath:  module.AfterScriptPath,
			PackagePath:      module.PackagePath,
		}

		mList = append(mList, newM)
	}

	m.Modules = mList
	return m
}

func FromTemplateModel(m *models.Template) *Template {
	t := &Template{
		ID:       m.ID,
		Name:     m.Name,
		Language: m.Language,
	}

	mList := make([]Module, 0, len(m.Modules))

	for _, module := range m.Modules {
		newM := Module{
			Index:            module.Index,
			Name:             module.Name,
			BeforeScriptPath: module.BeforeScriptPath,
			AfterScriptPath:  module.AfterScriptPath,
			PackagePath:      module.PackagePath,
		}

		mList = append(mList, newM)
	}

	t.Modules = mList
	return t
}
