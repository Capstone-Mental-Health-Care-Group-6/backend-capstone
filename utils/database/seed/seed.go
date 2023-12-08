package seed

import "gorm.io/gorm"

type Seed struct {
	Name string
	Run  func(*gorm.DB) error // No need to specify the type here
}
