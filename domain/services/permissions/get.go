package services

import (
	"SimonBK_Login/domain/models"
	"SimonBK_Login/infra/db"
	"fmt"
)

func GetResourcesAccessDetail() ([]models.ResourceRoleAccess, error) {
	var results []models.ResourceRoleAccess

	err := db.DBConn.
		Table("permissions_by_roles").
		Select("resources.name as resource, actions.action, roles.role_description").
		Joins("INNER JOIN permissions ON permissions_by_roles.fk_permission = permissions.id").
		Joins("INNER JOIN resources ON permissions.fk_resource = resources.id").
		Joins("INNER JOIN actions ON permissions.fk_action = actions.id").
		Joins("INNER JOIN roles ON permissions_by_roles.fk_role = roles.id").
		Scan(&results).Error

	if err != nil {
		return results, err
	}
	fmt.Println(results)

	return results, nil
}
