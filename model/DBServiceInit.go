package model

import "gorm.io/gorm"

type IServiceDb interface {
	CreateRestaurant(Restaurants) (RestaurantResponse, error)
	AddMenuItems(int, MenuItems) (MenuItem, error)
	GetRestaurant(int) (Restaurant, error)
	GetAllRestaurants() ([]Restaurant, error)
	GetAllMenuItems(int) ([]MenuItem, error)
}

type ServiceDb struct {
	DB *gorm.DB
}


func NewServiceDb(db *gorm.DB) *ServiceDb {
	return &ServiceDb{DB: db}
}

// mockgen -source=model/DBServiceInit.go -destination=service_db_mock.go catalogService/model IServiceDb