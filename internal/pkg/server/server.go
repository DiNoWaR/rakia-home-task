package server

import (
	"fmt"
	"github.com/dinowar/rakia-home-task/internal/pkg/domain/model"
	"github.com/dinowar/rakia-home-task/internal/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type CreatePostRequest struct {
	Author  string `json:"author" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content" binding:"required"`
}

type Server struct {
	rep    *service.RepositoryService
	logger *service.LogService
}

func NewAppServer(rep *service.RepositoryService, logger *service.LogService) *Server {
	return &Server{
		rep:    rep,
		logger: logger,
	}
}

func (server *Server) CreatePost(context *gin.Context) {
	var req CreatePostRequest
	if bindErr := context.ShouldBindJSON(&req); bindErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + bindErr.Error(),
		})
		return
	}

	post := &model.Post{
		Id:      uuid.NewString(),
		Author:  req.Author,
		Title:   req.Title,
		Content: req.Content,
	}

	dbPost := server.rep.CreatePost(post)
	context.JSON(http.StatusOK, gin.H{
		"id":      dbPost.Id,
		"author":  dbPost.Author,
		"content": dbPost.Content,
		"title":   dbPost.Title,
	})

}

func (server *Server) GetPosts(context *gin.Context) {
	author := context.DefaultQuery("author", "")
	posts := server.rep.GetPosts(author)
	context.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func (server *Server) GetPost(context *gin.Context) {
	postId := context.Param("id")
	post := server.rep.GetPost(postId)

	if post == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Post %s not found", postId),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func (server *Server) UpdatePost(context *gin.Context) {
	var req UpdatePostRequest
	if bindErr := context.ShouldBindJSON(&req); bindErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body: " + bindErr.Error(),
		})
		return
	}

	postId := context.Param("id")

	post := server.rep.UpdatePost(postId, req.Title, req.Content)
	if post == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Post %s not found", postId),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func (server *Server) DeletePost(context *gin.Context) {
	postId := context.Param("id")
	post := server.rep.DeletePost(postId)

	if post == nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("Post %s not found", postId),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}
