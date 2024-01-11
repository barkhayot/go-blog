package initializer

import "go-blog/model"

func SyncDatabase() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Post{})
}
