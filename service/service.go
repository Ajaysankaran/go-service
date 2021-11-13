package service

import "github.com/theAimOne/service-gateway/user"

type Service struct {
	Id            uint
	Name          string
	Subtitle      string
	Description   string
	User          user.User `gorm:"foreignKey:ID"`
	CategoryId    uint
	SubCategoryId uint
	Address       string
	Latitude      int
	Longitude     int
	Cost          float64
	Measure       string
}
