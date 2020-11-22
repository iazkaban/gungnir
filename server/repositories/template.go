package repositories

import "gungnir/entities"

type TemplateStorage interface {
	Add(*entities.Template) error
	Delete(id int64) error
	Update(*entities.Template) error
	GetAll() ([]*entities.Template, error)
	Get(id string) (*entities.Template, error)
}
