package models

//	type Resource struct {
//		gorm.Model
//		name        string `gorm:"column:name"`
//		description string `gorm:"column:description"`
//		path        string `gorm:"column:path"`
//	}
//
// Representación del resultado del query para obtener los privilegios de un rol
type ResourceRoleAccess struct {
	ResourceName string `gorm:"column:resource"`
	Action       string
	RoleDesc     string `gorm:"column:role_description"`
}
