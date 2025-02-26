package config

import (
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

// ImportCitizenshipData Import Citizenship Data is a function to import citizenship data from SQL file
func ImportCitizenshipData(db *gorm.DB) {
	fileName := "./config/SQLscript.sql"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		panic("SQL file is not found")
	}

	sqlScript, _ := os.ReadFile(fileName)
	sqlArr := strings.Split(string(sqlScript), ";")
	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}

		err := db.Exec(sql).Error
		if err != nil {
			panic("Error in importing SQL file")
		}
	}
	log.Println("Import SQL file successfully")
}
