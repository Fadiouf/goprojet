package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connexion à la base de donnée
func psql_connexion() {
	db_connexion_info := "host=localhost user=postgres password=Projet dbname=test port=5432 sslmode=prefer"
	db_connexion, err := gorm.Open(postgres.Open(db_connexion_info), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Migrate the schema
	db_connexion.AutoMigrate(&User{}, &AuthToken{}, &RefreshToken{}, &Role{}, &Group{})

}
