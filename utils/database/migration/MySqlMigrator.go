package migration

import (
	"gorm.io/gorm"
)

type MySqlMigrator struct {
	DB *gorm.DB
}

func NewMySqlMigrator(db *gorm.DB) Migrator {
	return &MySqlMigrator{
		DB: db,
	}
}

func (m *MySqlMigrator) CreateTable(table ...Table) {
	db := m.DB
	for _, t := range table {
		db.AutoMigrate(t)
	}
}

func (m *MySqlMigrator) DropTable(table ...Table) {
	db := m.DB
	for _, t := range table {
		db.Migrator().DropTable(t)
	}
}
