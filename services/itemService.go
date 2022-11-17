package services

import "online_store/models"

type ItemService interface {
    GetAll() ([]*models.Item, error)
    Get(*string) (*models.Item, error)
	CreateItem(*models.Item) error
    UpdateItem(*models.Item) error
    DeleteItem(*string) error
}