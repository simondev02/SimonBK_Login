// models/types.go
package swagger

import (
	"time"
)

// GormModelStub es una representación simplificada de gorm.Model para la documentación de Swagger.
type GormModelStub struct {
	ID        uint       `json:"id" example:"1"`
	CreatedAt time.Time  `json:"created_at" example:"2023-10-05T15:04:05Z"`
	UpdatedAt time.Time  `json:"updated_at" example:"2023-10-05T15:04:05Z"`
	DeletedAt *time.Time `json:"deleted_at" example:"2023-10-05T15:04:05Z"`
}
