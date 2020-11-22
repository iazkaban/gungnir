package impl

import (
	"errors"
	"fmt"
	"gungnir/entities"
	log "gungnir/logs"
	"gungnir/repositories"

	jsoniter "github.com/json-iterator/go"

	"xorm.io/xorm"
)

type TemplateStorageImpl struct {
	conn *xorm.Engine
}

func NewTemplateStorageImpl() (repositories.TemplateStorage, error) {
	conn, err := getDatabaseConn()
	if err != nil {
		log.Debugln(err)
		return nil, err
	}
	return &TemplateStorageImpl{
		conn: conn,
	}, nil
}

func (ts *TemplateStorageImpl) Add(template *entities.Template) error {
	_, err := ts.conn.InsertOne(template)
	fmt.Println(jsoniter.MarshalToString(template))
	return err
}

func (ts *TemplateStorageImpl) Delete(id int64) error {
	_, err := ts.conn.Where("id = ?", id).Delete(&entities.Template{})
	return err
}

func (ts *TemplateStorageImpl) Update(template *entities.Template) error {
	_, err := ts.conn.Update(template)
	return err
}

func (ts *TemplateStorageImpl) GetAll() ([]*entities.Template, error) {
	var list []*entities.Template
	err := ts.conn.OrderBy("id DESC").Find(&list)
	return list, err
}

func (ts *TemplateStorageImpl) Get(id string) (*entities.Template, error) {
	var t *entities.Template
	b, err := ts.conn.Where("id = ?", id).Get(t)
	if !b {
		return nil, errors.New("not found")
	}
	return t, err
}
