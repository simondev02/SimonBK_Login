package services

import (
	"SimonBK_Login/domain/models"
	"SimonBK_Login/infra/db"
)

func GetResourcesAccessDetail() ([]models.ResourceRoleAccess, error) {
	var results []models.ResourceRoleAccess

	err := db.DBConn.
		Table("permissions_by_roles").
		Select("resources.name as resource, permissions.permission, roles.\"RoleDescription\"").
		Joins("INNER JOIN permissions_by_resources on permissions_by_roles.\"Fk_PermissionsByResources\" = permissions_by_resources.id").
		Joins("INNER JOIN resources ON permissions_by_resources.\"Fk_Resource\" = resources.id").
		Joins("INNER JOIN permissions ON permissions_by_resources.\"Fk_Permission\" = permissions.id").
		Joins("INNER JOIN roles ON permissions_by_roles.\"Fk_Role\" = roles.id").
		Scan(&results).Error

	if err != nil {
		return results, err
	}

	return results, nil
}
