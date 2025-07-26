package database

import (
	"countreez/configs"
	"embed"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*.sql
var dbMigrations embed.FS

func DBMigrate() {
	migrations := &migrate.EmbedFileSystemMigrationSource{
       FileSystem: dbMigrations,
       Root:       "migrations",
    }

    n, errs := migrate.Exec(configs.DB, "postgres", migrations, migrate.Up)
    if errs != nil {
       panic(errs)
    }

    fmt.Println("Migration success, applied", n, "migrations!")
}