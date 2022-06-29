package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/xvbnm48/blogGO/post"
)

type postHandler struct {
	postService post.Service
}

func NewPostHandler(postService post.Service) *postHandler {
	return &postHandler{postService}
}

func (h *postHandler) FindAllPost(c *gin.Context) {
	post, err := h.postService.FindAll()
	if err != nil {

		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "find all post success",
		"post":    post,
	})
}

func (h *postHandler) FindPostById(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	post, err := h.postService.FindById(idInt)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "find post by id error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "find post by id success",
		"post":    post,
	})
}

func (h *postHandler) CreatePost(c *gin.Context) {
	var postRequest post.PostRequest
	err := c.BindJSON(&postRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s: , condition %s", e.Field(), e.Tag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(400, gin.H{
			"message": errorMessages,
		})
	}

	post, err := h.postService.Create(postRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "create post error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "create post success",
		"post":    post,
	})
}

func (h *postHandler) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	idString, _ := strconv.Atoi(id)
	var postRequest post.PostRequest
	err := c.ShouldBindJSON(&postRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field %s: , condition %s", e.Field(), e.Tag())
			errorMessages = append(errorMessages, errorMessage)
		}
	}

	post, err := h.postService.Update(idString, postRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "update post error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "update post success",
		"post":    post,
	})

}

func (h *postHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	idString, _ := strconv.Atoi(id)

	delete, err := h.postService.Delete(idString)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "delete post error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "delete post success",
		"delete":  delete,
	})
}
