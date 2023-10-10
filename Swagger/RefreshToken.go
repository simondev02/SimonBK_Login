package swagger

import "time"

type RefreshToken struct {
	Token           string
	FkUser          uint
	ExpiryDate      time.Time
	DeletedByUserID *uint
	UpdatedByUserID *uint
}
