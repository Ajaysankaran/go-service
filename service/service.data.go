package service

import (
	"log"

	"github.com/theAimOne/service-gateway/config"
)

func GetServices() ([]Service, error) {
	var service []Service
	result := config.DB.Find(&service)
	if result.Error != nil {
		log.Fatal(result.Error)
		return nil, result.Error
	}
	return service, nil
}

func CreateService(service Service) (uint, error) {
	result := config.DB.Create(&service)
	if result.Error != nil {
		log.Fatal(result.Error)
		return 0, result.Error
	}
	return service.Id, nil
}
