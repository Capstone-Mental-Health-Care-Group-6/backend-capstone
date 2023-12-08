package seeds

import (
	"FinalProject/utils/database/seed"

	"gorm.io/gorm"
)

func All() []seed.Seed {
	var seeds []seed.Seed = []seed.Seed{
		{
			Name: "CreateArticleCategory1",
			Run: func(db *gorm.DB) error {
				return CreateArticleCategory(db, "Anxiety")
			},
		},
		{
			Name: "CreateArticleCategory2",
			Run: func(db *gorm.DB) error {
				return CreateArticleCategory(db, "Depresi")
			},
		},
		{
			Name: "CreateArticleCategory3",
			Run: func(db *gorm.DB) error {
				return CreateArticleCategory(db, "Emosi")
			},
		},
		{
			Name: "CreateArticleCategory4",
			Run: func(db *gorm.DB) error {
				return CreateArticleCategory(db, "Kecemasan")
			},
		},
		{
			Name: "CreateArticleCategory5",
			Run: func(db *gorm.DB) error {
				return CreateArticleCategory(db, "Stress")
			},
		},
		{
			Name: "CreateArticleCategory6",
			Run: func(db *gorm.DB) error {
				return CreateArticleCategory(db, "Tips")
			},
		},
		{
			Name: "CreateArticleCategory7",
			Run: func(db *gorm.DB) error {
				return CreateArticleCategory(db, "Umum")
			},
		},
	}
	return seeds
}
