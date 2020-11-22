package repositories

import "gungnir/entities"

type ModuleStorage interface {
	Add(*entities.Module) error
	Delete(id int64) error
	DeleteByTemplateID(id int64) error
	GetByTemplateID(id int64) ([]*entities.Module, error)
	Get(id int64) (*entities.Module, error)
}
