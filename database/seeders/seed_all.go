package seeders

import (
	"log"
)

var seeders = []func(){
	SeedUsers,
}

func SeedAll() {
	log.Println("🚀 Starting database seeding...")

	for _, seedFunc := range seeders {
		seedFunc()
	}

	log.Println("✅ Database seeding completed!")
}
