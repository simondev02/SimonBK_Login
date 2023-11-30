package services

import (
	"SimonBK_Login/api/views"
	"SimonBK_Login/domain/models"
	services "SimonBK_Login/domain/services/permissions"
	"log"
)

func GetResourcesMap() ([]views.ResourceResponse, error) {
	// Realizar la consulta
	results, err := services.GetResourcesAccessDetail()
	if err != nil {
		log.Println("Error al consultar los detalles de los recursos:", err)
		return nil, err
	}

	// Transformar los resultados
	var resourcesMap = make(map[string]map[string]models.ActionRoles)
	for _, result := range results {
		if _, ok := resourcesMap[result.ResourceName]; !ok {
			resourcesMap[result.ResourceName] = make(map[string]models.ActionRoles)
		}
		actionRoles := resourcesMap[result.ResourceName][result.Permission]
		actionRoles.Roles = append(actionRoles.Roles, result.RoleDesc)
		resourcesMap[result.ResourceName][result.Permission] = actionRoles
	}

	var resources []views.ResourceResponse
	for name, actions := range resourcesMap {
		resources = append(resources, views.ResourceResponse{Resource: name, Actions: actions})
	}

	return resources, nil
}
