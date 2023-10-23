package models

import "SimonBK_Login/db"

type Permission struct {
	ID         int
	Permission string
}

type Resource struct {
	resource string
	actions  map[string]Permission
}

// ResourceResponse represents the structure of the response for resource retrieval
type ResourceResponse struct {
	ResourceName string                 `json:"resourceName"`
	Actions      map[string]interface{} `json:"actions"`
}

type TempResult struct {
	ResourceName string `gorm:"column:resource"`
	Permission   string
}

func GetResourcesByRole(roleID int) ([]ResourceResponse, error) {
	// Realizar la consulta
	var results []TempResult

	err := db.DBConn.
		Table("permissions_by_roles").
		Select("resources.name as resource, permissions.permission").
		Joins("INNER JOIN permissions_by_resources on permissions_by_roles.\"Fk_PermissionsByResources\" = permissions_by_resources.id").
		Joins("INNER JOIN resources ON permissions_by_resources.\"Fk_Resource\" = resources.id").
		Joins("INNER JOIN permissions ON permissions_by_resources.\"Fk_Permission\" = permissions.id").
		Where("permissions_by_roles.\"Fk_Role\" = ?", roleID).
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	// Transformar los resultados
	var resourcesMap = make(map[string]map[string]interface{})

	for _, result := range results {
		if _, ok := resourcesMap[result.ResourceName]; !ok {
			resourcesMap[result.ResourceName] = make(map[string]interface{})
		}
		resourcesMap[result.ResourceName][result.Permission] = struct{}{}
	}

	var resources []ResourceResponse
	for name, actions := range resourcesMap {
		resources = append(resources, ResourceResponse{ResourceName: name, Actions: actions})
	}

	return resources, nil
}
