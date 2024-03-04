package model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type MenuItems struct {
	gorm.Model
	ID           int    `json:"ID" gorm:"primaryKey"`
	Name         string  `json:"name" gorm:"not null"`
	Price        float32 `json:"price"`
	RestaurantsID int    `json:"restaurantsID"`
}

type MenuItem struct {
	ID           int    `json:"ID"`
	Name         string  `json:"name"`
	Price        float32 `json:"price"`
	RestaurantsID int    `json:"restaurantsID"`
}


func (s *ServiceDb) AddMenuItems(id int,item MenuItems) (MenuItem, error) {

	
	var restaurant Restaurants
	res := s.DB.First(&restaurant, id)
	if res.Error != nil {
		return MenuItem{}, fmt.Errorf("restaurant with ID %d not found", id)
	}

	restaurant.Menu = append(restaurant.Menu, item)


	item.RestaurantsID = restaurant.ID

	res1 := s.DB.Create(&item)
	if res1.RowsAffected == 0 {
		return MenuItem{}, errors.New("menu item creation unsuccessful")
	}

	convertedItem := MenuItem{
		ID: item.ID,
		Name: item.Name,
		Price: item.Price,
		RestaurantsID: item.RestaurantsID,
	}

	return convertedItem, nil
}


func (s *ServiceDb) GetAllMenuItems(userId int) ([]MenuItem, error) {

	var menuItems []MenuItems

	result := s.DB.Where("restaurants_id = ?", userId).Find(&menuItems)
	if result.Error != nil {
		return nil, result.Error
	}

	return Convert(menuItems), nil	
}


