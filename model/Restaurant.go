package model

import (
	"errors"

	"gorm.io/gorm"
)


type Restaurants struct{
	gorm.Model
	Name string
	Location string
	Menu     []MenuItems
}


func(s *ServiceDb) CreateRestaurant(restaurant Restaurants)(RestaurantResponse,error){
	
	
	res := s.DB.Create(&restaurant)
	if res.RowsAffected == 0{
		return RestaurantResponse{}, errors.New("restaurant creation unsuccessful")
	}

	response := RestaurantResponse{
		Message: "Restaurant created successfuly",
		Restaurant: restaurant,
	}

	return response,nil
}
