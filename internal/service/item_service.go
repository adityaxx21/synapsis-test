package service

import (
	"synapsis-backend-test/internal/domain"
	"synapsis-backend-test/config"

)


func CreateItem(item *domain.Item) error {
	if err := config.DB.Create(&item).Error; err != nil {
		return err
	}

	return nil
}


func UpdateItem(id string, item *domain.Item)  (*domain.Item, error) {
	var data *domain.Item

	if err :=  config.DB.Model(&item).Where("id = ?", id).Updates(item).Error; err != nil {
		return nil, err
	}

	if err := config.DB.First(&data, id).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func DeleteItem(id string, item *domain.Item) error {
	if err := config.DB.Delete(&item, id).Error; err != nil {
		return err
	}

	return nil
}

func DetailItem(id string)  (*domain.Item, error) {
	var item *domain.Item

	if err := config.DB.First(&item, id).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func ListItem(category string, page int, limit int)   ([]*domain.Item, error) {
	var items []*domain.Item

	if err := config.DB.Scopes(domain.FilterByCategory(category), domain.Paginate(page, limit)).Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}