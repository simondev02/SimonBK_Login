package models

import "SimonBK_Login/infra/db"

// ResourceResponse representa la estructura de la respuesta para la recuperación de recursos
type ResourceResponse struct {
	Resource string                 `json:"resource"`
	Actions  map[string]ActionRoles `json:"actions"`
}

type ActionRoles struct {
	Roles []string `json:"roles"`
}

type TempResult struct {
	ResourceName string `gorm:"column:resource"`
	Permission   string
	RoleDesc     string `gorm:"column:RoleDescription"`
}

func GetResourcesByRole() ([]ResourceResponse, error) {
	// Realizar la consulta
	var results []TempResult

	err := db.DBConn.
		Table("permissions_by_roles").
		Select("resources.name as resource, permissions.permission, roles.\"RoleDescription\"").
		Joins("INNER JOIN permissions_by_resources on permissions_by_roles.\"Fk_PermissionsByResources\" = permissions_by_resources.id").
		Joins("INNER JOIN resources ON permissions_by_resources.\"Fk_Resource\" = resources.id").
		Joins("INNER JOIN permissions ON permissions_by_resources.\"Fk_Permission\" = permissions.id").
		Joins("INNER JOIN roles ON permissions_by_roles.\"Fk_Role\" = roles.id").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	// Transformar los resultados
	var resourcesMap = make(map[string]map[string]ActionRoles)

	for _, result := range results {
		if _, ok := resourcesMap[result.ResourceName]; !ok {
			resourcesMap[result.ResourceName] = make(map[string]ActionRoles)
		}
		actionRoles := resourcesMap[result.ResourceName][result.Permission]
		actionRoles.Roles = append(actionRoles.Roles, result.RoleDesc)
		resourcesMap[result.ResourceName][result.Permission] = actionRoles
	}

	var resources []ResourceResponse
	for name, actions := range resourcesMap {
		resources = append(resources, ResourceResponse{Resource: name, Actions: actions})
	}

	return resources, nil
}
