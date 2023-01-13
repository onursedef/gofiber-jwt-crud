package postrepository

import (
	"backend.go/src/database"
	"backend.go/src/models/post"
)

func GetAll() ([]post.Post, error) {
	var posts []post.Post
	var err error

	if err := database.Connection().Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, err
}

func GetOne(id string) (post.Post, error) {
	var post post.Post
	var err error

	if err := database.Connection().First(&post, id).Error; err != nil {
		return post, err
	}
	return post, err
}

func Create(post post.Post) (post.Post, error) {
	var err error

	if err := database.Connection().Create(&post).Error; err != nil {
		return post, err
	}
	return post, err
}

func Update(post post.Post) (post.Post, error) {
	var err error

	if err := database.Connection().Save(&post).Error; err != nil {
		return post, err
	}
	return post, err
}

func Patch(id string, post post.Post) (post.Post, error) {
	var err error

	if err := database.Connection().Model(&post).Where("id = ?", id).Updates(post).Error; err != nil {
		return post, err
	}
	return post, err
}

func Delete(post post.Post) (post.Post, error) {
	var err error

	if err := database.Connection().Delete(&post).Error; err != nil {
		return post, err
	}
	return post, err
}

func DeleteAll() (bool, error) {
	var err error

	if err := database.Connection().Unscoped().Delete(&post.Post{}).Error; err != nil {
		return false, err
	}
	return true, err
}

func GetByTitle(title string) (post.Post, error) {
	var post post.Post
	var err error

	if err := database.Connection().Where("title = ?", title).First(&post).Error; err != nil {
		return post, err
	}
	return post, err
}

func GetByTitleLike(title string) ([]post.Post, error) {
	var posts []post.Post
	var err error

	if err := database.Connection().Where("title LIKE ?", "%"+title+"%").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, err
}

func GetBySlug(slug string) (post.Post, error) {
	var post post.Post
	var err error

	if err := database.Connection().Where("slug = ?", slug).First(&post).Error; err != nil {
		return post, err
	}
	return post, err
}
