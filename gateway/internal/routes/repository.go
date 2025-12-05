package routes

import (
	"gateway/internal/db"
	"gateway/internal/models"
)

func CreateRoute(route *models.Route) error {
	return db.GetDB().Create(route).Error
}

func GetRouteID(id uint) (*models.Route, error) {
	var route models.Route
	err := db.GetDB().First(&route, id).Error
	return &route, err
}

func ListRoutes() ([]models.Route, error) {
	var routes []models.Route
	err := db.GetDB().Find(&routes).Error
	return routes, err
}

func UpdateRoute(route *models.Route) error {
	return db.GetDB().Save(route).Error
}

func DeleteRoute(id uint) error {
	return db.GetDB().Delete(&models.Route{}, id).Error
}
