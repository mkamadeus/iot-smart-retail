package item

import "github.com/mkamadeus/iot-smart-retail/models"

func (s *Service) GetAll() ([]*models.ItemDAO, error) {
	items := make([]*models.Item, 0)
	result := s.Database.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}

	daos := make([]*models.ItemDAO, len(items))
	for i, item := range items {
		daos[i] = &models.ItemDAO{
			ID:   item.ID,
			Name: item.Name,
		}
	}

	return daos, nil
}
