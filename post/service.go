package post

import "errors"

type Service interface {
	FindAll() ([]Post, error)
	FindById(id int) (Post, error)
	Create(postrequest PostRequest) (Post, error)
	Update(id int, postrequest PostRequest) (Post, error)
	Delete(id int) (Post, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Post, error) {
	posts, err := s.repository.FindAllPost()
	return posts, err
}

func (s *service) FindById(id int) (Post, error) {
	post, err := s.repository.FindPostById(id)
	return post, err
}

func (s *service) Create(postrequest PostRequest) (Post, error) {
	post, err := s.repository.CreatePost(Post{
		Title:   postrequest.Title,
		Content: postrequest.Content,
	})
	return post, err
}

func (s *service) Update(id int, postrequest PostRequest) (Post, error) {
	post, _ := s.repository.FindPostById(id)
	if post.Id == 0 {
		return post, errors.New("post not found")
	}
	post.Title = postrequest.Title
	post.Content = postrequest.Content
	post, err := s.repository.Update(post)
	if err != nil {
		return post, err
	}

	newPost, err := s.repository.Update(post)
	if err != nil {
		return post, err
	}
	return newPost, nil
}

func (s *service) Delete(id int) (Post, error) {
	post, _ := s.repository.FindPostById(id)
	if post.Id == 0 {
		return post, errors.New("post not found")
	}

	deletePost, err := s.repository.Delete(post)
	if err != nil {
		return post, err
	}
	return deletePost, nil
}
