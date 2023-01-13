package migration

import (
	"backend.go/src/database"
	"backend.go/src/models/post"
	"backend.go/src/models/user"
)

func Migrate() {
	// Migration of post model
	database.Connection().AutoMigrate(&post.Post{})

	// Migration of user model
	database.Connection().AutoMigrate(&user.User{})
}
