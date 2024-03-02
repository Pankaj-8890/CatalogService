package model

import "gorm.io/gorm"

type IServiceDb interface {
	CreateRestaurant(Restaurants) (RestaurantResponse, error)
	AddMenuItems(int, MenuItems) (MenuItem, error)
	GetRestaurant(int) (Restaurant, error)
	GetAllRestaurants() ([]Restaurant, error)
	GetMenuItems(int) (MenuItems, error)
}

type ServiceDb struct {
	DB *gorm.DB
}


func NewServiceDb(db *gorm.DB) *ServiceDb {
	return &ServiceDb{DB: db}
}
