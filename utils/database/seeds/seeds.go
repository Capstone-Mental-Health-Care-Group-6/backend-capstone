package seeds

import (
	"FinalProject/utils/database/seed"

	"gorm.io/gorm"
)

func All() []seed.Seed {
	var seeds []seed.Seed = []seed.Seed{
		{
			Name: "CreateAdmin",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "Admin", "admin@gmail.com", "password", "Admin", "Active")
			},
		},
		{
			Name: "CreateUser",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "Irvan Hauwerich", "irvanhau@gmail.com", "password", "Doctor", "Active")
			},
		},
		{
			Name: "CreateArticleCategory1",
			Run: func(db *gorm.DB) error {
				return CreateArticleCategory(db, "Kecemasan", "slug")
			},
		},
	}
	return seeds
}
