package configs

import (
	"log"
	"os"

	"root/controllers"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "Singhasan26!",
		Addr:     "localhost:8080",
		Database: "postgres",
	}
	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to Connect")
		os.Exit(100)

	}
	log.Printf("CONNECTED TO DATABASE")
	controllers.CreateTenantTable(db)
	controllers.CreateTenantActivityTable(db)
	controllers.InitiateDB(db)
	return db
}
