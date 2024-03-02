package model

import "gorm.io/gorm"

type MenuItems struct {
	gorm.Model
	Name          string
	Price         float32
	RestaurantsID uint
}
