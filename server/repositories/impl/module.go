package impl

import (
	"gungnir/entities"
	log "gungnir/logs"
	"gungnir/repositories"

	"xorm.io/xorm"
)

type ModuleStorageImpl struct {
	conn *xorm.Engine
}

func NewModuleStorageImpl() (repositories.ModuleStorage, error) {
	conn, err := getDatabaseConn()
	if err != nil {
		log.Debugln(err)
		return nil, err
	}
	return &ModuleStorageImpl{
		conn: conn,
	}, nil
}

func (ms *ModuleStorageImpl) Add(module *entities.Module) error {
	_, err := ms.conn.InsertOne(module)
	return err
}

func (ms *ModuleStorageImpl) Delete(id int64) error {
	panic("implement me")
}

func (ms *ModuleStorageImpl) DeleteByTemplateID(id int64) error {
	_, err := ms.conn.Where("template_id = ?", id).Delete(&entities.Module{})
	return err
}

func (ms *ModuleStorageImpl) GetByTemplateID(id int64) ([]*entities.Module, error) {
	panic("implement me")
}

func (ms *ModuleStorageImpl) Get(id int64) (*entities.Module, error) {
	panic("implement me")
}
