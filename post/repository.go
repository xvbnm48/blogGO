package post

import "gorm.io/gorm"

type Repository interface {
	FindAllPost() ([]Post, error)
	FindPostById(id int) (Post, error)
	CreatePost(post Post) (Post, error)
	Update(post Post) (Post, error)
	Delete(post Post) (Post, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindAllPost() ([]Post, error) {
	var post []Post
	err := r.db.Find(&post).Error
	return post, err
}

func (r *repository) FindPostById(id int) (Post, error) {
	var post Post
	err := r.db.Where("id = ?", id).First(&post).Error
	return post, err
}

func (r *repository) CreatePost(post Post) (Post, error) {
	err := r.db.Create(&post).Error
	return post, err
}
func (r *repository) Update(post Post) (Post, error) {
	err := r.db.Save(&post).Error
	return post, err
}

func (r *repository) Delete(post Post) (Post, error) {
	err := r.db.Delete(&post).Error
	return post, err
}
