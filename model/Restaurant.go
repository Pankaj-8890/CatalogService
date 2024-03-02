package model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)
 

type Restaurants struct {
	gorm.Model
	ID       uint        `json:"ID" gorm:"primaryKey"`
	Name     string      `json:"name" gorm:"not null"`
	Location string      `json:"location" gorm:"not null"`
	Menu     []MenuItems `json:"menu" gorm:"foreignkey:RestaurantsID"`
}

type Restaurant struct {
	ID       uint        `json:"ID"`
	Name     string      `json:"name"`
	Location string      `json:"location"`
	Menu     []MenuItem `json:"menu"`
}


func (s *ServiceDb) CreateRestaurant(restaurant Restaurants) (RestaurantResponse, error) {

	res := s.DB.Create(&restaurant)
	if res.RowsAffected == 0 {
		return RestaurantResponse{}, errors.New("restaurant creation unsuccessful")
	}

	rest := Restaurant{
		ID:       restaurant.ID,
		Name:     restaurant.Name,
		Location: restaurant.Location,
		Menu:     []MenuItem{},
	}

	response := RestaurantResponse{
		Message:    "Restaurant created successfuly",
		Restaurant: rest,
	}

	return response, nil
}

func (s *ServiceDb) GetRestaurant(id int) (Restaurant, error) {


	var restaurant Restaurants
	res := s.DB.Preload("Menu").First(&restaurant, id)

	if res.Error != nil {
		return Restaurant{}, fmt.Errorf("restaurant with ID %d not found", id)
	}


	rest := Restaurant{
		ID:       restaurant.ID,
		Name:     restaurant.Name,
		Location: restaurant.Location,
		Menu:     convert(restaurant.Menu),
	}

	if res.Error != nil {
		return rest, fmt.Errorf("restaurant with ID %d not found", id)
	}

	return rest, nil
}

func (s *ServiceDb) GetAllRestaurants()([]Restaurant, error){

	var restaurants []Restaurants
	res := s.DB.Preload("Menu").Find(&restaurants)

	if res.Error != nil {
		return nil, res.Error
	}


	


	var convertedRestaurants []Restaurant
	for _, r := range restaurants {
		convertedRestaurants = append(convertedRestaurants, Restaurant{
			ID:       r.ID,
			Name:     r.Name,
			Location: r.Location,
			Menu:     convert(r.Menu),
		})
	}

	return convertedRestaurants, nil

}

func convert(menu []MenuItems)([]MenuItem){

	var menuItem []MenuItem
	for _,m := range menu{
		menuItem = append(menuItem, MenuItem{
			ID: m.ID,
			Name: m.Name,
			Price: m.Price,
			RestaurantsID: m.RestaurantsID,
		})
	}
	return menuItem
}