package services

import (
	"gungnir/entities"
	log "gungnir/logs"
	"gungnir/models"
	"gungnir/repositories/impl"
)

func AddTemplate(m *models.Template) (*models.Template, error) {
	repository, err := impl.NewTemplateStorageImpl()
	if err != nil {
		return nil, err
	}

	tempEntity := entities.FromTemplateModel(m)
	log.Debugln(tempEntity)
	err = repository.Add(tempEntity)
	if err != nil {
		return nil, err
	}

	moduleRepository, err := impl.NewModuleStorageImpl()
	if err != nil {
		return nil, err
	}

	for _, module := range m.Modules {
		err = moduleRepository.Add(entities.FromModuleModel(&module, tempEntity.ID))
		log.Errorln(err)
	}
	return tempEntity.ToModel(), nil
}

func ListTemplate() ([]*models.Template, error) {
	repository, err := impl.NewTemplateStorageImpl()
	if err != nil {
		return nil, err
	}

	list, err := repository.GetAll()
	if err != nil {
		return nil, err
	}

	mList := make([]*models.Template, len(list))

	for k, v := range list {
		mList[k] = v.ToModel()
	}

	return mList, nil
}

func DeleteTemplate(id int64) error {
	repository, err := impl.NewTemplateStorageImpl()
	if err != nil {
		return err
	}

	err = repository.Delete(id)
	if err != nil {
		return err
	}

	moduleRepository, err := impl.NewModuleStorageImpl()
	if err != nil {
		return err
	}

	err = moduleRepository.DeleteByTemplateID(id)
	if err != nil {
		return err
	}
	return nil
}
