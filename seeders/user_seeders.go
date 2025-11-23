package seeders

import (
	"Test/config"
	"Test/models"
	"Test/utilities/general"
	"log"

	"gorm.io/gorm"
)

var seedUsersData = []models.User{
	{Username: "admin", Email: "admin@example.com", Password: "123456"},
	{Username: "test", Email: "test@example.com", Password: "password"},
}

func SeedUsers() {
	for _, user := range seedUsersData {
		var userExists models.User
		err := config.DB.Where("email = ? OR username = ?", user.Email, user.Username).First(&userExists).Error
		if err == nil {
			// کاربر موجود است → ادامه بدون اضافه کردن
			log.Printf("⚠️ User already exists: %s", user.Username)
			continue
		}
		if err != nil && err != gorm.ErrRecordNotFound {
			log.Printf("❌ Error checking user %s: %v", user.Username, err)
			continue
		}
		user.Password = general.MakeHash(user.Password)

		if err := config.DB.Create(&user).Error; err != nil {
			log.Printf("❌ Error adding user %s: %v", user.Username, err)
		} else {
			log.Printf("✅ User added successfully: %s", user.Username)
		}
	}
}
