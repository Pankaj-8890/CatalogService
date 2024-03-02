package model

import "gorm.io/gorm"

type IServiceDb interface {
	CreateRestaurant(Restaurants) (RestaurantResponse,error) 
	// AddMenuItems()
	// GetRestaurant()
	// GetMenuItems()
}

type ServiceDb struct {
	DB *gorm.DB
}


func NewServiceDb(db *gorm.DB) *ServiceDb {
	return &ServiceDb{DB: db}	
}


