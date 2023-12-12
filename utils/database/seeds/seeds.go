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
		{
			Name: "CreateCounselingMethod1",
			Run: func(db *gorm.DB) error {
				return CreateCounselingMethod(db, "Chat", 20000)
			},
		},
		{
			Name: "CreateCounselingMethod2",
			Run: func(db *gorm.DB) error {
				return CreateCounselingMethod(db, "VideoCall", 40000)
			},
		},
		{
			Name: "CreateCounselingDuration1",
			Run: func(db *gorm.DB) error {
				return CreateCounselingDuration(db, "60 Minute", 10000)
			},
		},
		{
			Name: "CreateCounselingDuration2",
			Run: func(db *gorm.DB) error {
				return CreateCounselingDuration(db, "90 Minute", 20000)
			},
		},
		{
			Name: "CreateCounselingTopic1",
			Run: func(db *gorm.DB) error {
				return CreateCounselingTopic(db, "Pekerjaan")
			},
		},
		{
			Name: "CreateCounselingTopic2",
			Run: func(db *gorm.DB) error {
				return CreateCounselingTopic(db, "Kendali Emosi")
			},
		},
		{
			Name: "CreateCounselingTopic3",
			Run: func(db *gorm.DB) error {
				return CreateCounselingTopic(db, "Percintaan")
			},
		},
		{
			Name: "CreateCounselingTopic4",
			Run: func(db *gorm.DB) error {
				return CreateCounselingTopic(db, "Pendidikan")
			},
		},
		{
			Name: "CreateCounselingTopic5",
			Run: func(db *gorm.DB) error {
				return CreateCounselingTopic(db, "Keluarga")
			},
		},
		{
			Name: "CreateCounselingTopic6",
			Run: func(db *gorm.DB) error {
				return CreateCounselingTopic(db, "Kecanduan")
			},
		},
		{
			Name: "CreateCounselingTopic7",
			Run: func(db *gorm.DB) error {
				return CreateCounselingTopic(db, "Kesepian")
			},
		},
		{
			Name: "CreateCounselingTopic8",
			Run: func(db *gorm.DB) error {
				return CreateCounselingTopic(db, "Sosial")
			},
		},
		{
			Name: "CreateCounselingTopic9",
			Run: func(db *gorm.DB) error {
				return CreateCounselingTopic(db, "Lainnya")
			},
		},
	}
	return seeds
}
