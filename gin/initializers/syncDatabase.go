package initializers

import "myapp/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Post{})
}
