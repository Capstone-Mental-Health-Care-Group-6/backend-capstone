package migration

type Migrator interface {
	CreateTable(table ...Table)
	DropTable(table ...Table)
}

type Table interface {
	TableName() string
}
