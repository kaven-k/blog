package services

import (
	"server/services/image_service"
	"server/services/user_service"
)

type ServiceGroup struct {
	ImageService image_service.ImageService
	UserService  user_service.UserService
}

var ServiceApp = new(ServiceGroup)
