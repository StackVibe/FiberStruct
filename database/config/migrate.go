package config

import (
	"Test/core/models"
	"fmt"
)

func MigrateModels() {
	modelsList := []interface{}{
		&models.User{},
	}

	for _, model := range modelsList {
		err := DB.AutoMigrate(model)
		if err != nil {
			fmt.Println("❌ Error To Migrate Models : ", err)
		} else {
			fmt.Println("✅ Migration Success : ", model)
		}
	}
}
